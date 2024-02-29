package stocks

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"htmx/pkg/app"
	loglib "htmx/pkg/log"
	"htmx/pkg/otel"
	ot "htmx/pkg/otel"
	"htmx/pkg/otel/log"
	"htmx/pkg/server"
	"os"
	"os/signal"
)

type Stocks struct {
	Ctx           context.Context
	config        app.Config
	stop          context.CancelFunc
	shutdownFuncs []func(context.Context) error
}

var _ app.Lifecycle = &Stocks{}

func NewApp(config app.Config) *Stocks {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	return &Stocks{
		Ctx:    ctx,
		config: config,
		stop:   stop,
	}
}

func (app *Stocks) Setup() (app.ShutdownFunc, error) {
	log.SetLogger(loglib.NewStdLogger(app.config.Name, loglib.DEBUG, loglib.NewFileExporter()))

	log.Debug(app.Ctx, "Setting up application")

	// Set up OpenTelemetry.
	log.Debug(app.Ctx, "Setting up OpenTelemetry")
	otelShutdown, err := ot.SetupOTelSDK(app.Ctx, app.config.Name, app.config.Version)
	if err != nil {
		return nil, err
	}

	return otelShutdown, nil
}

func (app *Stocks) Run() (app.ShutdownFunc, error) {
	// Setup HTTP server.
	log.Debug(app.Ctx, "Starting HTTP server")
	srv, err := server.NewServer(server.Config{
		Addr:   ":420",
		Routes: app.routes(),
	})
	if err != nil {
		log.Fatal(app.Ctx, "unable to create server", otel.ErrorMsg.String(err.Error()))
	}

	// Start HTTP server.
	log.Debug(app.Ctx, "Running application")
	srvErr := make(chan error, 1)
	go func() {
		srvErr <- srv.ListenAndServe()
	}()

	// Wait
	select {
	case <-srvErr:
		// Error when starting HTTP server.
		log.Error(app.Ctx, "Error starting HTTP server")
	case <-app.Ctx.Done():
		// Wait for first CTRL+C.
		log.Fatal(app.Ctx, "Context Cancelled")
	}

	// Stop HTTP server.
	log.Info(app.Ctx, "Shutting down")
	shutdown := func(ctx context.Context) error {
		return srv.Shutdown(ctx)
	}
	return shutdown, nil
}

func (app *Stocks) Cleanup(shutdownFuncs []app.ShutdownFunc) {
	log.Info(app.Ctx, "Cleaning up application")
	for _, fn := range shutdownFuncs {
		log.Info(app.Ctx, "cleanup error", otel.ErrorMsg.String(fmt.Sprintf("%s", fn(app.Ctx))))
	}
	app.stop()
}

func connectToPostgreSQL() (*gorm.DB, error) {
	dsn := "user=myuser password=mypassword dbname=mydatabase host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

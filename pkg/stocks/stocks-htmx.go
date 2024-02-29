package stocks

import (
	"context"
	"fmt"
	"htmx/pkg/app"
	loglib "htmx/pkg/log"
	"htmx/pkg/otel"
	ot "htmx/pkg/otel"
	"htmx/pkg/otel/log"
	"htmx/pkg/server"
	"htmx/pkg/stocks/internal/pgdb"
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
	conn := fmt.Sprintf("host=localhost user=%s dbname=%s sslmode=disable password=%s", app.config.DbUser, app.config.DbName, app.config.DbPassword)
	if err := pgdb.Connect(conn); err != nil {
		return nil, fmt.Errorf("Failed to connect to database:", err)
	}

	if err := pgdb.Migrate(); err != nil {
		return nil, fmt.Errorf("Failed to migrate database:", err)
	}

	// pgdb.Seed()

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
		Addr:   "localhost:4200",
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
	case err := <-srvErr:
		log.Error(app.Ctx, "Error starting HTTP server", otel.ErrorMsg.String(err.Error()))
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

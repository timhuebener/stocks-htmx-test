package stocks

import (
	"context"
	"fmt"
	"htmx/pkg/app"
	loglib "htmx/pkg/log"
	"htmx/pkg/otel"
	ot "htmx/pkg/otel"
	"htmx/pkg/otel/log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Stocks struct {
	Ctx           context.Context
	config        app.Config
	stop          context.CancelFunc
	shutdownFuncs []func(context.Context) error
}

var _ app.Lifecycle = (*Stocks)(nil)

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

	// db, err := connectToPostgreSQL()
	// if err != nil {
	// 	log.Println(err)
	// }

	// // Perform database migration
	// err = db.AutoMigrate(&Product{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return otelShutdown, nil
}

func (app *Stocks) Run() (app.ShutdownFunc, error) {
	// Setup HTTP server.
	log.Debug(app.Ctx, "Starting HTTP server")
	srv := &http.Server{
		Addr:         ":8080",
		BaseContext:  func(_ net.Listener) context.Context { return app.Ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      newHTTPHandler(),
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

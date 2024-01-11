package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	handlers "htmx/pkg/handlers"
	ot "htmx/pkg/otel"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	if err := application(); err != nil {
		log.Fatalln(err)
	}
}

func application() (err error) {
	log.Println("Starting application")

	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	log.Println("Setting up OpenTelemetry")
	serviceName := "htmx"
	serviceVersion := "0.1.0"
	otelShutdown, err := ot.SetupOTelSDK(ctx, serviceName, serviceVersion)
	if err != nil {
		log.Println("Error setting up OpenTelemetry")
		return err
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		log.Println("Shutting down OpenTelemetry")
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	// Start HTTP server.
	log.Println("Starting HTTP server")
	srv := &http.Server{
		Addr:         ":8080",
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      newHTTPHandler(),
	}
	srvErr := make(chan error, 1)
	go func() {
		srvErr <- srv.ListenAndServe()
	}()

	// Wait for interruption.
	log.Println("Waiting for interruption")
	select {
	case err = <-srvErr:
		// Error when starting HTTP server.
		log.Fatalln(err)
	case <-ctx.Done():
		// Wait for first CTRL+C.
		stop()
	}

	log.Println("Shutting down")

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	return srv.Shutdown(ctx)
}

func newHTTPHandler() http.Handler {
	mux := http.NewServeMux()

	// handleFunc is a replacement for mux.HandleFunc
	// which enriches the handler's HTTP instrumentation with the pattern as the http.route.
	handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
		// Configure the "http.route" for the HTTP instrumentation.
		handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		mux.Handle(pattern, handler)
	}

	// Register handlers.
	handleFunc("/", handlers.Root)
	handleFunc("/hello", handlers.Hello)

	// Add HTTP instrumentation for the whole server.
	handler := otelhttp.NewHandler(mux, "/")
	return handler
}

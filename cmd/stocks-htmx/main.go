package main

import (
	"htmx/pkg/app"
	"htmx/pkg/otel"
	"htmx/pkg/otel/log"
	"htmx/pkg/stocks"
)

func main() {
	config := app.Config{
		Name:    "stocks-htmx",
		Version: "0.0.1",
	}

	stocks := stocks.NewApp(config)

	var shutdownFuncs []app.ShutdownFunc

	shutdown, err := stocks.Setup()
	if err != nil {
		log.Error(stocks.Ctx, "setup error", otel.ErrorMsg.String(err.Error()))
	}
	shutdownFuncs = append(shutdownFuncs, shutdown)

	shutdown, err = stocks.Run()
	if err != nil {
		log.Error(stocks.Ctx, "run error", otel.ErrorMsg.String(err.Error()))
	}
	shutdownFuncs = append(shutdownFuncs, shutdown)

	stocks.Cleanup(shutdownFuncs)
}

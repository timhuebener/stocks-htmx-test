package main

import (
	"fmt"
	"htmx/pkg/app"
	"htmx/pkg/otel"
	"htmx/pkg/otel/log"
	"htmx/pkg/stocks"
	"os"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	config := app.Config{
		Name:       "stocks-htmx",
		Version:    "0.0.1",
		BasePath:   cwd + "/web/templates",
		DbName:     "mydatabase",
		DbUser:     "myuser",
		DbPassword: "mypassword",
	}

	stocks := stocks.NewApp(config)

	var shutdownFuncs []app.ShutdownFunc

	shutdown, err := stocks.Setup()
	if err != nil {
		log.Error(stocks.Ctx, "setup error", otel.ErrorMsg.String(err.Error()))
		return
	}
	shutdownFuncs = append(shutdownFuncs, shutdown)

	shutdown, err = stocks.Run()
	if err != nil {
		log.Error(stocks.Ctx, "run error", otel.ErrorMsg.String(err.Error()))
		return
	}
	shutdownFuncs = append(shutdownFuncs, shutdown)

	stocks.Cleanup(shutdownFuncs)
}

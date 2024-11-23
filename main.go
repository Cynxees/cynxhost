package main

import (
	"context"
	"cynxhost/internal/app"
	"cynxhost/internal/controller"
	"flag"

	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	envFile := flag.String("env", ".env", ".env")
	flag.Parse()

	// Load the specified .env file
	err := godotenv.Load(*envFile)
	if err != nil {
		panic(err)
	}

	app, err := app.NewApp(ctx, "config.json")
	if err != nil {
		panic(err)
	}

	logger := app.Dependencies.Logger

	logger.Infoln("Creating http server")
	httpServer, err := controller.NewHttpServer(app);
	if err != nil {
		panic(err)
	}

	logger.Infoln("Starting http server")
	if err := httpServer.Start(); err != nil {
		panic(err)
	}

	<-ctx.Done()
}
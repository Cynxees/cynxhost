package main

import (
	"context"
	"mchost/internal/app"
	"mchost/internal/controller"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

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
package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/kupriyanovkk/gophkeeper/internal/server/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cancel()

	app, err := app.NewApp(ctx)
	if err != nil {
		panic(err)
	}

	app.Server.Start(cancel)

	<-ctx.Done()
	app.Server.Stop()

	app.Logger.Info("Server stopped")
}

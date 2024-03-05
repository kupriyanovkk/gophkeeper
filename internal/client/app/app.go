package app

import "context"

type App struct {
}

func NewApp(ctx context.Context) (*App, error) {
	return &App{}, nil
}

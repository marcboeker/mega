package app

import (
	"github.com/marcboeker/mega/config"
	"github.com/marcboeker/mega/db"
	"github.com/marcboeker/mega/server"
)

type App struct {
	cfg    *config.Config
	client *db.Client
	srv    *server.Server
}

func New(cfg *config.Config, client *db.Client, srv *server.Server) *App {
	return &App{cfg, client, srv}
}

func (a *App) Run() error {
	return a.srv.Start()
}

func (a *App) Shutdown() error {
	if err := a.client.Close(); err != nil {
		return err
	}
	return a.srv.Close()
}

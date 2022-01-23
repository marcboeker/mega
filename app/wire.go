//go:build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/marcboeker/mega/config"
	"github.com/marcboeker/mega/db"
	"github.com/marcboeker/mega/logger"
	"github.com/marcboeker/mega/server"
	"github.com/marcboeker/mega/service"
)

func Initialize() (*App, error) {
	wire.Build(config.New, db.New, service.New, logger.New, server.New, New)
	return &App{}, nil
}

package server

import (
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/marcboeker/mega/config"
	"github.com/marcboeker/mega/graph/resolver"
	"github.com/marcboeker/mega/service"
)

type Server struct {
	*echo.Echo
	cfg *config.Config
}

func New(cfg *config.Config, services *service.Mapper, logger *log.Logger) *Server {
	s := Server{cfg: cfg}
	s.Echo = echo.New()
	srv := handler.NewDefaultServer(resolver.NewSchema(services, logger))
	{
		queryPath := cfg.GetString("server.queryPath")
		s.Echo.POST(queryPath, func(c echo.Context) error {
			srv.ServeHTTP(c.Response(), c.Request())
			return nil
		})

		if cfg.GetBool("server.enablePlayground") {
			s.Echo.GET(cfg.GetString("server.playgroundPath"), func(c echo.Context) error {
				playground.Handler("GraphQL", queryPath).ServeHTTP(c.Response(), c.Request())
				return nil
			})
		}
	}

	s.Echo.Use(injectLogger(logger))

	if cfg.GetBool("server.logging") {
		s.Echo.Use(middleware.Logger())
	}

	if cfg.GetBool("server.recoveryMiddleware") {
		s.Echo.Use(middleware.Recover())
	}

	return &s
}

func (s *Server) Start() error {
	return s.Echo.Start(s.cfg.GetString("server.address"))
}

func (s *Server) Close() error {
	return s.Echo.Close()
}

func injectLogger(logger *log.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// c.Set("logger", )
			c.SetLogger(c.Logger())
			return next(c)
		}
	}
}

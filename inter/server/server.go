package server

import (
	"avitotech-api/cfg"
	"avitotech-api/inter/cases"
	"avitotech-api/inter/handlers"
	_ "github.com/docker/docker/daemon/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	ucase *cases.UC
	config  *cfg.Config
}

func NewServer(cfg cfg.Config, ucase *cases.UC) *Server {
	return &Server{
		ucase: ucase,
		config:  &cfg,
	}
}

func (s *Server) Start() error {
	app := fiber.New()
	app.Use(logger.New())
	api := app.Group("/api")

	reg := handlers.Reg{Ucase: s.ucase}
	reg.Registeration(api)

	return app.Listen(s.config.Server.Address)
}

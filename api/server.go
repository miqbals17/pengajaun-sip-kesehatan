package api

import "github.com/gofiber/fiber/v3"

type Server struct {
	router *fiber.App
}

func NewServer() (*Server, error) {
	server := &Server{}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/users", func(c fiber.Ctx) error {
		return c.SendString("List of all users!")
	})

	server.router = app
}

func (server *Server) Start(address string) error {
	return server.router.Listen(address)
}

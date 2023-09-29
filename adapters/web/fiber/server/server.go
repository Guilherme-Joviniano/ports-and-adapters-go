package server

import (
	"log"
	"time"

	"github.com/Guilherme-Joviniano/go-hexagonal/adapters/web/fiber/handler"
	"github.com/Guilherme-Joviniano/go-hexagonal/application/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type WebServer struct {
	Service domain.ProductServiceInterface
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {

	server := fiber.New(fiber.Config{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	server.Use(logger.New())

	handler.MakeProductHandler(server, w.Service)

	err := server.Listen(":9000")

	if err != nil {
		log.Fatal(err.Error())
	}
}

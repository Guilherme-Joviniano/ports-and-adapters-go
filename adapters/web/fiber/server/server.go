package server

import (
	"log"
	"os"

	routes "github.com/Guilherme-Joviniano/go-hexagonal/adapters/web/fiber/routes/private"
	"github.com/Guilherme-Joviniano/go-hexagonal/application/main/factories"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var address = os.Getenv("HTTP_ADDRESS")

type WebServer struct {
	server *fiber.App
}

func MakeNewWebServer() *WebServer {
	server := fiber.New()
	return &WebServer{
		server: server,
	}
}

func (w *WebServer) Serve() {
	w.server.Use(logger.New())

	w.SetRoutes()

	err := w.server.Listen(address)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func (w *WebServer) SetRoutes() {
	productService := factories.MakeProductService()
	routes.ProductRouter(w.server, productService)
}

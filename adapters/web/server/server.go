package server

import (
	"log"
	"time"

	"github.com/Guilherme-Joviniano/go-hexagonal/application/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/urfave/negroni"
)

type WebServer struct {
	Service domain.ProductServiceInterface
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	n := negroni.New(
		negroni.NewLogger(),
	)

	server := fiber.New(fiber.Config{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	err := server.Server().ListenAndServe(":8080")

	if err != nil {
		log.Fatal(err.Error())
	}
}

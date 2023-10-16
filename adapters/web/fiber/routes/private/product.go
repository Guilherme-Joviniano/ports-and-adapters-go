package routes

import (
	"github.com/Guilherme-Joviniano/go-hexagonal/adapters/web/fiber/handler"
	"github.com/Guilherme-Joviniano/go-hexagonal/application/domain"
	"github.com/gofiber/fiber/v2"
)

func ProductRouter(router *fiber.App, service domain.ProductServiceInterface) {
	router.Post("/products", handler.CreateProductHandler(service))
	router.Get("/products/:id", handler.GetProductHandler(service))
	router.Patch("/products/enable/:id", handler.EnableProductHandler(service))
	router.Patch("/products/disable/:id", handler.DisableProductHandler(service))
}

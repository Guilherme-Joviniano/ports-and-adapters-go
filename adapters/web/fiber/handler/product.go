package handler

import (
	web_utils "github.com/Guilherme-Joviniano/go-hexagonal/adapters/web/fiber/utils"
	"github.com/Guilherme-Joviniano/go-hexagonal/application/domain"
	"github.com/Guilherme-Joviniano/go-hexagonal/application/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func MakeProductHandler(router *fiber.App, service domain.ProductServiceInterface) {
	router.Post("/products", createProductHandler(service))
	router.Get("/products/:id", getProductHandler(service))
	router.Patch("/products/enable/:id", enableProductHandler(service))
	router.Patch("/products/disable/:id", disableProductHandler(service))
}

func disableProductHandler(service domain.ProductServiceInterface) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Accepts("application/json")
		id := ctx.Params("id")

		product, err := service.Get(id)

		if err != nil {
			return web_utils.NotFound(ctx, err)
		}

		result, err := service.Disable(product)

		if err != nil {
			return web_utils.UnprocessableEntity(ctx, err)
		}

		return ctx.JSON(result)
	}
}

func enableProductHandler(service domain.ProductServiceInterface) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Accepts("application/json")
		id := ctx.Params("id")

		product, err := service.Get(id)

		if err != nil {
			return web_utils.NotFound(ctx, err)
		}

		result, err := service.Enable(product)

		if err != nil {
			return web_utils.UnprocessableEntity(ctx, err)
		}

		return ctx.JSON(result)
	}
}

func createProductHandler(service domain.ProductServiceInterface) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		p := dto.NewProduct()

		if err := ctx.BodyParser(p); err != nil {
			return web_utils.UnprocessableEntity(ctx, err)
		}

		response, err := service.Create(p.Name, p.Price)

		if err != nil {
			return web_utils.UnprocessableEntity(ctx, err)
		}

		return ctx.JSON(response)
	}
}

func getProductHandler(service domain.ProductServiceInterface) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Accepts("application/json")
		id := ctx.Params("id")

		product, err := service.Get(id)

		if err != nil {
			return web_utils.NotFound(ctx, err)
		}

		return ctx.JSON(product)
	}
}

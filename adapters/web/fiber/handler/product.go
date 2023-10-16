package handler

import (
	web_utils "github.com/Guilherme-Joviniano/go-hexagonal/adapters/web/fiber/utils"
	"github.com/Guilherme-Joviniano/go-hexagonal/application/domain"
	"github.com/Guilherme-Joviniano/go-hexagonal/application/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func DisableProductHandler(service domain.ProductServiceInterface) fiber.Handler {
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

func EnableProductHandler(service domain.ProductServiceInterface) fiber.Handler {
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

func CreateProductHandler(service domain.ProductServiceInterface) fiber.Handler {
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

func GetProductHandler(service domain.ProductServiceInterface) fiber.Handler {
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

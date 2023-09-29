package web_utils

import "github.com/gofiber/fiber/v2"

func UnprocessableEntity(ctx *fiber.Ctx, err error) error {
	ctx.SendStatus(422)
	return ctx.JSON(NewErrorJson(err.Error(), 422))
}

func NotFound(ctx *fiber.Ctx, err error) error {
	ctx.SendStatus(404)
	return ctx.JSON(NewErrorJson(err.Error(), 404))
}

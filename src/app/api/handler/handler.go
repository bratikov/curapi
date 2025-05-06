package handler

import (
	"currency/app/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Init(app *fiber.App) {
	main := app.Group("/", logger.New())
	main.Get("/", ping)
}

func ping(c *fiber.Ctx) error {
	return c.Status(fiber.StatusAccepted).JSON(api.ResponseSuccess(api.DataEmpty{}))
}

package app

import (
	"currency/app/api/handler"
	"currency/app/repository"
	"currency/app/service"
	"currency/internal/config"
	"currency/internal/logs"
	"fmt"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "currency/docs"
)

var (
	server *fiber.App
)

// @title			Currency API
// @version		1.0
// @description	Currency API documentation
// @contact.name	Alexander Bratikov
// @contact.email	alexander.bratikov@gmail.com
func Serve() {
	repository.Connect()
	defer repository.Close()

	service.Init()

	server := fiber.New(fiber.Config{
		AppName:      "Currency",
		ServerHeader: "Currency v1",
		BodyLimit:    128 * 1024 * 1024,
	})

	server.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logs.Logger,
	}))

	server.Get("/docs/*", swagger.HandlerDefault)

	handler.Init(server)

	err := server.Listen(fmt.Sprintf("%s:%d", config.Currency.Host, config.Currency.Port))
	if err != nil {
		logs.Fatal("Cant start pastor server", err)
	}
}

func Shutdown() {
	err := server.Shutdown()
	if err != nil {
		logs.Error("Shutdown server error", err)
	}
}

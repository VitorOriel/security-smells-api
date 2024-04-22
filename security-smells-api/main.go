package main

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
	"security-smells-api/routes"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 15 * 1024 * 1024, // this is the default limit of 15MB
	})
	logger, _ := zap.NewProduction()

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))
	log.Info("Starting server")

	routes.SetupRoutes(app)
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

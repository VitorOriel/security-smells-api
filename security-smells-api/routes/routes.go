package routes

import (
	"github.com/gofiber/fiber/v2"
	"security-smells-api/controller"
	"security-smells-api/repository"
	"security-smells-api/service"
)

func SetupRoutes(app *fiber.App) {
	smellyController := inject()
	app.Post("/api/v1/smelly", smellyController.Execute)
}

func inject() controller.SmellyController {
	return controller.SmellyController{
		SmellyService: service.SmellyService{
			SmellyRepository: repository.SmellyRepository{},
		},
	}
}

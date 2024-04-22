package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"security-smells-api/controller"
	"security-smells-api/models"
	"security-smells-api/repository"
	"security-smells-api/service"
)

var Validator = validator.New()

func ValidateSmelly(c *fiber.Ctx) error {
	var errors []*models.IError
	body := new(models.Smelly)
	c.BodyParser(&body)

	err := Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el models.IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	return c.Next()
}

func SetupRoutes(app *fiber.App) {
	smellyController := inject()
	app.Post("/api/v1/smelly", ValidateSmelly, smellyController.Execute)
}

func inject() controller.SmellyController {
	return controller.SmellyController{
		SmellyService: service.SmellyService{
			SmellyRepository: repository.SmellyRepository{},
		},
	}
}

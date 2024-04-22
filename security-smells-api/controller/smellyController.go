package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"security-smells-api/models"
	"security-smells-api/service"
)

type SmellyController struct {
	SmellyService service.SmellyService
}

func (smellyController SmellyController) Execute(c *fiber.Ctx) error {
	log.Info("Executing smelly controller")
	smelly := new(models.Smelly)
	if err := c.BodyParser(smelly); err != nil {
		return err
	}
	log.Info("Smelly received", smelly)
	return nil
}

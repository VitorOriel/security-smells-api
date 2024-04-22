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
	pods, deployments, statefulsets, daemonset, err := smellyController.SmellyService.Execute(smelly.YamlToValidate)
	if err != nil {
		return c.JSON(
			models.SmellyResponseErrorDTO{YamlToValidate: smelly.YamlToValidate, Message: err.Error()},
		)
	}
	log.Info("Pods", pods)
	log.Info("Deployments", deployments)
	log.Info("StatefulSets", statefulsets)
	log.Info("DaemonSets", daemonset)
	return nil
}

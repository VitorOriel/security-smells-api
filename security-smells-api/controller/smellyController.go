package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"security-smells-api/service"
)

type SmellyController struct {
	SmellyService service.SmellyService
}

func (smellyController SmellyController) Execute(c *fiber.Ctx) error {
	fmt.Println("Executing smelly controller")
	return nil
}

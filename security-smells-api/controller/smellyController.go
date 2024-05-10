package controller

import (
	"security-smells-api/models"
	"security-smells-api/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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
	pods, replicaSets, deployments, statefulsets, daemonset, jobs, cronJobs, err := smellyController.SmellyService.Execute(smelly.YamlToValidate)
	if err != nil {
		return c.JSON(
			models.SmellyResponseErrorDTO{YamlToValidate: smelly.YamlToValidate, Message: err.Error()},
		)
	}
	log.Info("Pods", pods)
	log.Info("ReplicaSets", replicaSets)
	log.Info("Deployments", deployments)
	log.Info("StatefulSets", statefulsets)
	log.Info("DaemonSets", daemonset)
	log.Info("Jobs", jobs)
	log.Info("CronJobs", cronJobs)

	smellsPod := smellyController.SmellyService.FindPodSmell(pods)
	smellsReplicaSet := smellyController.SmellyService.FindReplicaSetSmell(replicaSets)
	smellsDeployment := smellyController.SmellyService.FindDeploymentSmell(deployments)
	smellsJob := smellyController.SmellyService.FindJobSmell(jobs)
	smellsCronJob := smellyController.SmellyService.FindCronJobSmell(cronJobs)
	smellyResponseDTO := models.SmellyResponseDTO{
		TotalOfSmells:    len(smellsPod) + len(smellsReplicaSet) + len(smellsDeployment) + len(smellsJob) + len(smellsCronJob),
		SmellsPod:        smellsPod,
		SmellsReplicaSet: smellsReplicaSet,
		SmellsDeployment: smellsDeployment,
		SmellsJob:        smellsJob,
		SmellsCronJob:    smellsCronJob,
	}
	return c.JSON(smellyResponseDTO)
}

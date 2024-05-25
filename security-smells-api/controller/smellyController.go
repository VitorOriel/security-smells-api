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
	kubernetesWorkloads, err := smellyController.SmellyService.Execute(smelly.YamlToValidate)
	if err != nil {
		return c.JSON(
			models.SmellyResponseErrorDTO{YamlToValidate: smelly.YamlToValidate, Message: err.Error()},
		)
	}
	log.Info("Pods", kubernetesWorkloads.Pods)
	log.Info("ReplicaSets", kubernetesWorkloads.ReplicaSets)
	log.Info("Deployments", kubernetesWorkloads.Deployments)
	log.Info("StatefulSets", kubernetesWorkloads.StatefulSets)
	log.Info("DaemonSets", kubernetesWorkloads.DaemonSets)
	log.Info("Jobs", kubernetesWorkloads.Jobs)
	log.Info("CronJobs", kubernetesWorkloads.CronJobs)

	smellsPod := smellyController.SmellyService.FindPodSmell(kubernetesWorkloads.Pods)
	smellsReplicaSet := smellyController.SmellyService.FindReplicaSetSmell(kubernetesWorkloads.ReplicaSets)
	smellsDeployment := smellyController.SmellyService.FindDeploymentSmell(kubernetesWorkloads.Deployments)
	smellsJob := smellyController.SmellyService.FindJobSmell(kubernetesWorkloads.Jobs)
	smellsCronJob := smellyController.SmellyService.FindCronJobSmell(kubernetesWorkloads.CronJobs)
	smellsStatefulSet := smellyController.SmellyService.FindStatefulSetSmell(kubernetesWorkloads.StatefulSets)
	smellsDaemonSet := smellyController.SmellyService.FindDaemonSetSmell(kubernetesWorkloads.DaemonSets)
	smellyResponseDTO := models.SmellyResponseDTO{
		Meta: models.Meta{
			TotalOfSmells: len(smellsPod) + len(smellsReplicaSet) + len(smellsDeployment) + len(smellsJob) + len(smellsCronJob),
		},
		Data: models.Data{
			SmellsPod:         smellsPod,
			SmellsReplicaSet:  smellsReplicaSet,
			SmellsDeployment:  smellsDeployment,
			SmellsJob:         smellsJob,
			SmellsCronJob:     smellsCronJob,
			SmellsStatefulSet: smellsStatefulSet,
			SmellDemonSet:     smellsDaemonSet,
		},
	}
	return c.JSON(smellyResponseDTO)
}

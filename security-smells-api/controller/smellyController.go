package controller

import (
	"security-smells-api/constants"
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
	kubernetesWorkloads, fileMetadata, err := smellyController.SmellyService.Execute(smelly.YamlToValidate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			models.SmellyResponseErrorDTO{YamlToValidate: smelly.YamlToValidate, Message: err.Error()},
		)
	}
	log.Infof("Pods: %+v", kubernetesWorkloads.Pods)
	log.Infof("ReplicaSets: %+v", kubernetesWorkloads.ReplicaSets)
	log.Infof("Deployments: %+v", kubernetesWorkloads.Deployments)
	log.Infof("StatefulSets: %+v", kubernetesWorkloads.StatefulSets)
	log.Infof("DaemonSets: %+v", kubernetesWorkloads.DaemonSets)
	log.Infof("Jobs: %+v", kubernetesWorkloads.Jobs)
	log.Infof("CronJobs: %+v", kubernetesWorkloads.CronJobs)

	smellsPod := smellyController.SmellyService.FindPodSmell(kubernetesWorkloads.Pods, fileMetadata.WorkloadPositionMapByManifest[constants.POD_WORKLOAD])
	smellsReplicaSet := smellyController.SmellyService.FindReplicaSetSmell(kubernetesWorkloads.ReplicaSets, fileMetadata.WorkloadPositionMapByManifest[constants.REPLICASET_WORKLOAD])
	smellsDeployment := smellyController.SmellyService.FindDeploymentSmell(kubernetesWorkloads.Deployments, fileMetadata.WorkloadPositionMapByManifest[constants.DEPLOYMENT_WORKLOAD])
	smellsJob := smellyController.SmellyService.FindJobSmell(kubernetesWorkloads.Jobs, fileMetadata.WorkloadPositionMapByManifest[constants.JOB_WORKLOAD])
	smellsCronJob := smellyController.SmellyService.FindCronJobSmell(kubernetesWorkloads.CronJobs, fileMetadata.WorkloadPositionMapByManifest[constants.CRONJOB_WORKLOAD])
	smellsStatefulSet := smellyController.SmellyService.FindStatefulSetSmell(kubernetesWorkloads.StatefulSets, fileMetadata.WorkloadPositionMapByManifest[constants.STATEFULSET_WORKLOAD])
	smellsDaemonSet := smellyController.SmellyService.FindDaemonSetSmell(kubernetesWorkloads.DaemonSets, fileMetadata.WorkloadPositionMapByManifest[constants.DAEMONSET_WORKLOAD])
	smellyResponseDTO := models.SmellyResponseDTO{
		Meta: &models.Meta{
			TotalOfSmells:    len(smellsPod) + len(smellsReplicaSet) + len(smellsDeployment) + len(smellsJob) + len(smellsCronJob) + len(smellsStatefulSet) + len(smellsDaemonSet),
			DecodedWorkloads: models.NewWorkloadMeta(kubernetesWorkloads),
		},
		Data: &models.Data{
			SmellsPod:         smellsPod,
			SmellsReplicaSet:  smellsReplicaSet,
			SmellsDeployment:  smellsDeployment,
			SmellsJob:         smellsJob,
			SmellsCronJob:     smellsCronJob,
			SmellsStatefulSet: smellsStatefulSet,
			SmellsDaemonSet:   smellsDaemonSet,
		},
	}
	return c.JSON(smellyResponseDTO)
}

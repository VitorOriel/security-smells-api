package implementation

import (
	appsv1 "k8s.io/api/apps/v1"
	"security-smells-api/models"
	"security-smells-api/service/interfaces"
)

type Deployment struct {
	interfaces.SmellyDeployment
	Deployment      *appsv1.Deployment
	SmellDeployment []models.SmellDeployment
}

func (deployment *Deployment) SmellyResourceAndLimit() {
	// Check if the deployment has resource limits set
	//verify for all containers
	nameSpace := deployment.Deployment.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	deploymentName := deployment.Deployment.GetName()
	kind := deployment.Deployment.GroupVersionKind().Kind
	for _, container := range deployment.Deployment.Spec.Template.Spec.Containers {
		if container.Resources.Limits == nil {
			smellDeployment := models.SmellDeployment{
				NameSpace:      nameSpace,
				DeploymentName: deploymentName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "Resource limits not set for container " + container.Name,
				Suggestion:     "Set resource limits for container " + container.Name,
			}
			deployment.SmellDeployment = append(deployment.SmellDeployment, smellDeployment)
		}
		if container.Resources.Requests == nil {
			smellDeployment := models.SmellDeployment{
				NameSpace:      nameSpace,
				DeploymentName: deploymentName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "Resource requests not set for container " + container.Name,
				Suggestion:     "Set resource requests for container " + container.Name,
			}
			deployment.SmellDeployment = append(deployment.SmellDeployment, smellDeployment)
		}
	}
}

func (deployment *Deployment) SmellySecurityContext() {
	//TODO implement me
	panic("implement me")
}

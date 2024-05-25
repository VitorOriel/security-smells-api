package implementation

import (
	"security-smells-api/models"
	"security-smells-api/service/interfaces"

	appsv1 "k8s.io/api/apps/v1"
)

type Deployment struct {
	interfaces.SmellyDeployment
	Deployment       *appsv1.Deployment
	WorkloadPosition int
	SmellKubernetes  []models.SmellKubernetes
}

func (deployment *Deployment) SmellySecurityContextReadOnlyRootFilesystem() {
	nameSpace := deployment.Deployment.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	deploymentName := deployment.Deployment.GetName()
	kind := deployment.Deployment.GroupVersionKind().Kind
	for _, container := range deployment.Deployment.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.ReadOnlyRootFilesystem == nil {
			smellDeployment := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: deploymentName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "ReadOnlyRootFilesystem not set into " + container.Name + " your container is running with ReadWriteRootFilesystem",
				Suggestion:        "Please add ReadOnlyRootFilesystem into " + container.Name + " to avoid running with ReadWriteRootFilesystem",
			}
			deployment.SmellKubernetes = append(deployment.SmellKubernetes, smellDeployment)
		}
	}
}

func (deployment *Deployment) SmellySecurityContextAllowPrivilegeEscalation() {
	nameSpace := deployment.Deployment.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	deploymentName := deployment.Deployment.GetName()
	kind := deployment.Deployment.GroupVersionKind().Kind
	for _, container := range deployment.Deployment.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			smellDeployment := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: deploymentName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "AllowPrivilegeEscalation not set into " + container.Name + " your container is running with AllowPrivilegeEscalation",
				Suggestion:        "Please add AllowPrivilegeEscalation into " + container.Name + " to avoid running with AllowPrivilegeEscalation",
			}
			deployment.SmellKubernetes = append(deployment.SmellKubernetes, smellDeployment)
		}
	}
}

func (deployment *Deployment) SmellySecurityContextCapabilities() {
	nameSpace := deployment.Deployment.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	deploymentName := deployment.Deployment.GetName()
	kind := deployment.Deployment.GroupVersionKind().Kind
	for _, container := range deployment.Deployment.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			smellDeployment := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: deploymentName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "Capabilities not set into " + container.Name + " your container is running with full capabilities",
				Suggestion:        "Please add capabilities into " + container.Name + " to avoid running with full capabilities",
			}
			deployment.SmellKubernetes = append(deployment.SmellKubernetes, smellDeployment)
		}

	}
}

func (deployment *Deployment) SmellySecurityContextRunAsUser() {
	nameSpace := deployment.Deployment.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	deploymentName := deployment.Deployment.GetName()
	kind := deployment.Deployment.GroupVersionKind().Kind
	for _, container := range deployment.Deployment.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			smellDeployment := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: deploymentName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "RunAsUser not set into " + container.Name + " your container is running as root",
				Suggestion:        "Please add runAsUser into " + container.Name + " to avoid running as root",
			}
			deployment.SmellKubernetes = append(deployment.SmellKubernetes, smellDeployment)
		}
	}
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
			smellDeployment := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: deploymentName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "Resource limits not set for container " + container.Name,
				Suggestion:        "Set resource limits for container " + container.Name,
			}
			deployment.SmellKubernetes = append(deployment.SmellKubernetes, smellDeployment)
		}
		if container.Resources.Requests == nil {
			smellDeployment := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: deploymentName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "Resource requests not set for container " + container.Name,
				Suggestion:        "Set resource requests for container " + container.Name,
			}
			deployment.SmellKubernetes = append(deployment.SmellKubernetes, smellDeployment)
		}
	}
}

package implementation

import (
	"security-smells-api/constants"
	"security-smells-api/models"
	"security-smells-api/service/interfaces"

	appsv1 "k8s.io/api/apps/v1"
)

type Deployment struct {
	interfaces.SmellyDeployment
	Deployment       *appsv1.Deployment
	WorkloadPosition int
	SmellKubernetes  []*models.SmellKubernetes
}

func (deployment *Deployment) SmellySecurityContextReadOnlyRootFilesystem() {
	for _, container := range deployment.Deployment.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.ReadOnlyRootFilesystem == nil {
			deployment.SmellKubernetes = append(deployment.SmellKubernetes, models.NewSmellKubernetes(deployment.Deployment, deployment.Deployment.GetObjectKind(), &container, deployment.WorkloadPosition, constants.K8S_SEC_ROROOTFS))
		}
	}
}

func (deployment *Deployment) SmellySecurityContextAllowPrivilegeEscalation() {
	for _, container := range deployment.Deployment.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			deployment.SmellKubernetes = append(deployment.SmellKubernetes, models.NewSmellKubernetes(deployment.Deployment, deployment.Deployment.GetObjectKind(), &container, deployment.WorkloadPosition, constants.K8S_SEC_PRIVESCALATION))
		}
	}
}

func (deployment *Deployment) SmellySecurityContextCapabilities() {
	for _, container := range deployment.Deployment.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			deployment.SmellKubernetes = append(deployment.SmellKubernetes, models.NewSmellKubernetes(deployment.Deployment, deployment.Deployment.GetObjectKind(), &container, deployment.WorkloadPosition, constants.K8S_SEC_CAPABILITIES))
		}
	}
}

func (deployment *Deployment) SmellySecurityContextRunAsUser() {
	for _, container := range deployment.Deployment.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			deployment.SmellKubernetes = append(deployment.SmellKubernetes, models.NewSmellKubernetes(deployment.Deployment, deployment.Deployment.GetObjectKind(), &container, deployment.WorkloadPosition, constants.K8S_SEC_RUNASUSER))
		}
	}
}

func (deployment *Deployment) SmellyResourceAndLimit() {
	for _, container := range deployment.Deployment.Spec.Template.Spec.Containers {
		if container.Resources.Limits == nil {
			deployment.SmellKubernetes = append(deployment.SmellKubernetes, models.NewSmellKubernetes(deployment.Deployment, deployment.Deployment.GetObjectKind(), &container, deployment.WorkloadPosition, constants.K8S_SEC_RESLIMITS))
		}
		if container.Resources.Requests == nil {
			deployment.SmellKubernetes = append(deployment.SmellKubernetes, models.NewSmellKubernetes(deployment.Deployment, deployment.Deployment.GetObjectKind(), &container, deployment.WorkloadPosition, constants.K8S_SEC_RESREQUESTS))
		}
	}
}

func (deployment *Deployment) SmellySecurityContextPrivileged() {
	for _, container := range deployment.Deployment.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Privileged == nil {
			deployment.SmellKubernetes = append(deployment.SmellKubernetes, models.NewSmellKubernetes(deployment.Deployment, deployment.Deployment.GetObjectKind(), &container, deployment.WorkloadPosition, constants.K8S_SEC_PRIVILEGED))
		}
	}
}

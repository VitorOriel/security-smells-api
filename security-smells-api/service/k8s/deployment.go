package k8s

import (
	"security-smells-api/constants"
	"security-smells-api/models"

	appsv1 "k8s.io/api/apps/v1"
)

type Deployment struct {
	Deployment       *appsv1.Deployment
	WorkloadPosition int
	KubernetesSmell  []*models.KubernetesSmell
}

func (deployment *Deployment) SmellySecurityContextReadOnlyRootFilesystem() {
	for _, container := range deployment.Deployment.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.ReadOnlyRootFilesystem == nil {
			deployment.KubernetesSmell = append(deployment.KubernetesSmell, models.NewKubernetesSmell(deployment.Deployment, deployment.Deployment.GetObjectKind(), &container, deployment.WorkloadPosition, constants.K8S_SEC_ROROOTFS_UNSET))
		}
	}
}

func (deployment *Deployment) SmellySecurityContextAllowPrivilegeEscalation() {
	for _, container := range deployment.Deployment.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			deployment.KubernetesSmell = append(deployment.KubernetesSmell, models.NewKubernetesSmell(deployment.Deployment, deployment.Deployment.GetObjectKind(), &container, deployment.WorkloadPosition, constants.K8S_SEC_PRIVESCALATION_UNSET))
		}
	}
}

func (deployment *Deployment) SmellySecurityContextCapabilities() {
	for _, container := range deployment.Deployment.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			deployment.KubernetesSmell = append(deployment.KubernetesSmell, models.NewKubernetesSmell(deployment.Deployment, deployment.Deployment.GetObjectKind(), &container, deployment.WorkloadPosition, constants.K8S_SEC_CAPABILITIES_UNSET))
		}
	}
}

func (deployment *Deployment) SmellySecurityContextRunAsUser() {
	for _, container := range deployment.Deployment.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			deployment.KubernetesSmell = append(deployment.KubernetesSmell, models.NewKubernetesSmell(deployment.Deployment, deployment.Deployment.GetObjectKind(), &container, deployment.WorkloadPosition, constants.K8S_SEC_RUNASUSER_UNSET))
		}
	}
}

func (deployment *Deployment) SmellyResourceAndLimit() {
	for _, container := range deployment.Deployment.Spec.Template.Spec.Containers {
		if container.Resources.Limits == nil {
			deployment.KubernetesSmell = append(deployment.KubernetesSmell, models.NewKubernetesSmell(deployment.Deployment, deployment.Deployment.GetObjectKind(), &container, deployment.WorkloadPosition, constants.K8S_SEC_RESLIMITS_UNSET))
		}
		if container.Resources.Requests == nil {
			deployment.KubernetesSmell = append(deployment.KubernetesSmell, models.NewKubernetesSmell(deployment.Deployment, deployment.Deployment.GetObjectKind(), &container, deployment.WorkloadPosition, constants.K8S_SEC_RESREQUESTS_UNSET))
		}
	}
}

func (deployment *Deployment) SmellySecurityContextPrivileged() {
	for _, container := range deployment.Deployment.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Privileged == nil {
			deployment.KubernetesSmell = append(deployment.KubernetesSmell, models.NewKubernetesSmell(deployment.Deployment, deployment.Deployment.GetObjectKind(), &container, deployment.WorkloadPosition, constants.K8S_SEC_PRIVILEGED_UNSET))
		}
	}
}

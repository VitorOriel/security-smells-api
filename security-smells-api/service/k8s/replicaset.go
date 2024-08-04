package k8s

import (
	"security-smells-api/constants"
	"security-smells-api/models"

	appsv1 "k8s.io/api/apps/v1"
)

type ReplicaSet struct {
	ReplicaSet       *appsv1.ReplicaSet
	WorkloadPosition int
	KubernetesSmell  []*models.KubernetesSmell
}

func (replicaSet *ReplicaSet) SmellySecurityContextReadOnlyRootFilesystem() {
	for _, container := range replicaSet.ReplicaSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.ReadOnlyRootFilesystem == nil {
			replicaSet.KubernetesSmell = append(replicaSet.KubernetesSmell, models.NewKubernetesSmell(replicaSet.ReplicaSet, replicaSet.ReplicaSet.GetObjectKind(), &container, replicaSet.WorkloadPosition, constants.K8S_SEC_ROROOTFS_UNSET))
		}
	}
}

func (replicaSet *ReplicaSet) SmellySecurityContextAllowPrivilegeEscalation() {
	for _, container := range replicaSet.ReplicaSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			replicaSet.KubernetesSmell = append(replicaSet.KubernetesSmell, models.NewKubernetesSmell(replicaSet.ReplicaSet, replicaSet.ReplicaSet.GetObjectKind(), &container, replicaSet.WorkloadPosition, constants.K8S_SEC_PRIVESCALATION_UNSET))
		}
	}
}

func (replicaSet *ReplicaSet) SmellySecurityContextCapabilities() {
	for _, container := range replicaSet.ReplicaSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			replicaSet.KubernetesSmell = append(replicaSet.KubernetesSmell, models.NewKubernetesSmell(replicaSet.ReplicaSet, replicaSet.ReplicaSet.GetObjectKind(), &container, replicaSet.WorkloadPosition, constants.K8S_SEC_CAPABILITIES_UNSET))
		}
	}
}

func (replicaSet *ReplicaSet) SmellySecurityContextRunAsUser() {
	for _, container := range replicaSet.ReplicaSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			replicaSet.KubernetesSmell = append(replicaSet.KubernetesSmell, models.NewKubernetesSmell(replicaSet.ReplicaSet, replicaSet.ReplicaSet.GetObjectKind(), &container, replicaSet.WorkloadPosition, constants.K8S_SEC_RUNASUSER_UNSET))
		}
	}
}

func (replicaSet *ReplicaSet) SmellyResourceAndLimit() {
	for _, container := range replicaSet.ReplicaSet.Spec.Template.Spec.Containers {
		if container.Resources.Limits == nil {
			replicaSet.KubernetesSmell = append(replicaSet.KubernetesSmell, models.NewKubernetesSmell(replicaSet.ReplicaSet, replicaSet.ReplicaSet.GetObjectKind(), &container, replicaSet.WorkloadPosition, constants.K8S_SEC_RESLIMITS_UNSET))
		}
		if container.Resources.Requests == nil {
			replicaSet.KubernetesSmell = append(replicaSet.KubernetesSmell, models.NewKubernetesSmell(replicaSet.ReplicaSet, replicaSet.ReplicaSet.GetObjectKind(), &container, replicaSet.WorkloadPosition, constants.K8S_SEC_RESREQUESTS_UNSET))
		}
	}
}

func (replicaSet *ReplicaSet) SmellySecurityContextPrivileged() {
	for _, container := range replicaSet.ReplicaSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Privileged == nil {
			replicaSet.KubernetesSmell = append(replicaSet.KubernetesSmell, models.NewKubernetesSmell(replicaSet.ReplicaSet, replicaSet.ReplicaSet.GetObjectKind(), &container, replicaSet.WorkloadPosition, constants.K8S_SEC_PRIVILEGED_UNSET))
		}
	}
}

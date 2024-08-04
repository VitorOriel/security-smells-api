package implementation

import (
	"security-smells-api/constants"
	"security-smells-api/models"
	"security-smells-api/service/interfaces"

	appsv1 "k8s.io/api/apps/v1"
)

type ReplicaSet struct {
	interfaces.SmellyReplicaSet
	ReplicaSet       *appsv1.ReplicaSet
	WorkloadPosition int
	SmellKubernetes  []*models.SmellKubernetes
}

func (replicaSet *ReplicaSet) SmellySecurityContextReadOnlyRootFilesystem() {
	for _, container := range replicaSet.ReplicaSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.ReadOnlyRootFilesystem == nil {
			replicaSet.SmellKubernetes = append(replicaSet.SmellKubernetes, models.NewSmellKubernetes(replicaSet.ReplicaSet, replicaSet.ReplicaSet.GetObjectKind(), &container, replicaSet.WorkloadPosition, constants.K8S_SEC_ROROOTFS_UNSET))
		}
	}
}

func (replicaSet *ReplicaSet) SmellySecurityContextAllowPrivilegeEscalation() {
	for _, container := range replicaSet.ReplicaSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			replicaSet.SmellKubernetes = append(replicaSet.SmellKubernetes, models.NewSmellKubernetes(replicaSet.ReplicaSet, replicaSet.ReplicaSet.GetObjectKind(), &container, replicaSet.WorkloadPosition, constants.K8S_SEC_PRIVESCALATION_UNSET))
		}
	}
}

func (replicaSet *ReplicaSet) SmellySecurityContextCapabilities() {
	for _, container := range replicaSet.ReplicaSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			replicaSet.SmellKubernetes = append(replicaSet.SmellKubernetes, models.NewSmellKubernetes(replicaSet.ReplicaSet, replicaSet.ReplicaSet.GetObjectKind(), &container, replicaSet.WorkloadPosition, constants.K8S_SEC_CAPABILITIES_UNSET))
		}
	}
}

func (replicaSet *ReplicaSet) SmellySecurityContextRunAsUser() {
	for _, container := range replicaSet.ReplicaSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			replicaSet.SmellKubernetes = append(replicaSet.SmellKubernetes, models.NewSmellKubernetes(replicaSet.ReplicaSet, replicaSet.ReplicaSet.GetObjectKind(), &container, replicaSet.WorkloadPosition, constants.K8S_SEC_RUNASUSER_UNSET))
		}
	}
}

func (replicaSet *ReplicaSet) SmellyResourceAndLimit() {
	for _, container := range replicaSet.ReplicaSet.Spec.Template.Spec.Containers {
		if container.Resources.Limits == nil {
			replicaSet.SmellKubernetes = append(replicaSet.SmellKubernetes, models.NewSmellKubernetes(replicaSet.ReplicaSet, replicaSet.ReplicaSet.GetObjectKind(), &container, replicaSet.WorkloadPosition, constants.K8S_SEC_RESLIMITS_UNSET))
		}
		if container.Resources.Requests == nil {
			replicaSet.SmellKubernetes = append(replicaSet.SmellKubernetes, models.NewSmellKubernetes(replicaSet.ReplicaSet, replicaSet.ReplicaSet.GetObjectKind(), &container, replicaSet.WorkloadPosition, constants.K8S_SEC_RESREQUESTS_UNSET))
		}
	}
}

func (replicaSet *ReplicaSet) SmellySecurityContextPrivileged() {
	for _, container := range replicaSet.ReplicaSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Privileged == nil {
			replicaSet.SmellKubernetes = append(replicaSet.SmellKubernetes, models.NewSmellKubernetes(replicaSet.ReplicaSet, replicaSet.ReplicaSet.GetObjectKind(), &container, replicaSet.WorkloadPosition, constants.K8S_SEC_PRIVILEGED_UNSET))
		}
	}
}

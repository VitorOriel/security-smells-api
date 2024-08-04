package k8s

import (
	"security-smells-api/constants"
	"security-smells-api/models"

	v1 "k8s.io/api/apps/v1"
)

type StatefulSet struct {
	StatefulSet      *v1.StatefulSet
	WorkloadPosition int
	KubernetesSmell  []*models.KubernetesSmell
}

func (statefulSet *StatefulSet) SmellySecurityContextReadOnlyRootFilesystem() {
	for _, container := range statefulSet.StatefulSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.ReadOnlyRootFilesystem == nil {
			statefulSet.KubernetesSmell = append(statefulSet.KubernetesSmell, models.NewKubernetesSmell(statefulSet.StatefulSet, statefulSet.StatefulSet.GetObjectKind(), &container, statefulSet.WorkloadPosition, constants.K8S_SEC_ROROOTFS_UNSET))
		}
	}
}

func (statefulSet *StatefulSet) SmellySecurityContextAllowPrivilegeEscalation() {
	for _, container := range statefulSet.StatefulSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			statefulSet.KubernetesSmell = append(statefulSet.KubernetesSmell, models.NewKubernetesSmell(statefulSet.StatefulSet, statefulSet.StatefulSet.GetObjectKind(), &container, statefulSet.WorkloadPosition, constants.K8S_SEC_PRIVESCALATION_UNSET))
		}
	}
}

func (statefulSet *StatefulSet) SmellySecurityContextCapabilities() {
	for _, container := range statefulSet.StatefulSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			statefulSet.KubernetesSmell = append(statefulSet.KubernetesSmell, models.NewKubernetesSmell(statefulSet.StatefulSet, statefulSet.StatefulSet.GetObjectKind(), &container, statefulSet.WorkloadPosition, constants.K8S_SEC_CAPABILITIES_UNSET))
		}
	}
}

func (statefulSet *StatefulSet) SmellySecurityContextRunAsUser() {
	for _, container := range statefulSet.StatefulSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			statefulSet.KubernetesSmell = append(statefulSet.KubernetesSmell, models.NewKubernetesSmell(statefulSet.StatefulSet, statefulSet.StatefulSet.GetObjectKind(), &container, statefulSet.WorkloadPosition, constants.K8S_SEC_RUNASUSER_UNSET))
		}
	}
}

func (statefulSet *StatefulSet) SmellyResourceAndLimit() {
	for _, container := range statefulSet.StatefulSet.Spec.Template.Spec.Containers {
		if container.Resources.Requests == nil {
			statefulSet.KubernetesSmell = append(statefulSet.KubernetesSmell, models.NewKubernetesSmell(statefulSet.StatefulSet, statefulSet.StatefulSet.GetObjectKind(), &container, statefulSet.WorkloadPosition, constants.K8S_SEC_RESREQUESTS_UNSET))
		}
		if container.Resources.Limits == nil {
			statefulSet.KubernetesSmell = append(statefulSet.KubernetesSmell, models.NewKubernetesSmell(statefulSet.StatefulSet, statefulSet.StatefulSet.GetObjectKind(), &container, statefulSet.WorkloadPosition, constants.K8S_SEC_RESLIMITS_UNSET))
		}
	}
}

func (statefulSet *StatefulSet) SmellySecurityContextPrivileged() {
	for _, container := range statefulSet.StatefulSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Privileged == nil {
			statefulSet.KubernetesSmell = append(statefulSet.KubernetesSmell, models.NewKubernetesSmell(statefulSet.StatefulSet, statefulSet.StatefulSet.GetObjectKind(), &container, statefulSet.WorkloadPosition, constants.K8S_SEC_PRIVILEGED_UNSET))
		}
	}
}

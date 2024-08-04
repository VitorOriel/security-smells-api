package implementation

import (
	"security-smells-api/constants"
	"security-smells-api/models"
	"security-smells-api/service/interfaces"

	v1 "k8s.io/api/apps/v1"
)

type StatefulSet struct {
	interfaces.SmellyStatefulSet
	StatefulSet      *v1.StatefulSet
	WorkloadPosition int
	SmellKubernetes  []*models.SmellKubernetes
}

func (statefulSet *StatefulSet) SmellySecurityContextReadOnlyRootFilesystem() {
	for _, container := range statefulSet.StatefulSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.ReadOnlyRootFilesystem == nil {
			statefulSet.SmellKubernetes = append(statefulSet.SmellKubernetes, models.NewSmellKubernetes(statefulSet.StatefulSet, statefulSet.StatefulSet.GetObjectKind(), &container, statefulSet.WorkloadPosition, constants.K8S_SEC_ROROOTFS_UNSET))
		}
	}
}

func (statefulSet *StatefulSet) SmellySecurityContextAllowPrivilegeEscalation() {
	for _, container := range statefulSet.StatefulSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			statefulSet.SmellKubernetes = append(statefulSet.SmellKubernetes, models.NewSmellKubernetes(statefulSet.StatefulSet, statefulSet.StatefulSet.GetObjectKind(), &container, statefulSet.WorkloadPosition, constants.K8S_SEC_PRIVESCALATION_UNSET))
		}
	}
}

func (statefulSet *StatefulSet) SmellySecurityContextCapabilities() {
	for _, container := range statefulSet.StatefulSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			statefulSet.SmellKubernetes = append(statefulSet.SmellKubernetes, models.NewSmellKubernetes(statefulSet.StatefulSet, statefulSet.StatefulSet.GetObjectKind(), &container, statefulSet.WorkloadPosition, constants.K8S_SEC_CAPABILITIES_UNSET))
		}
	}
}

func (statefulSet *StatefulSet) SmellySecurityContextRunAsUser() {
	for _, container := range statefulSet.StatefulSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			statefulSet.SmellKubernetes = append(statefulSet.SmellKubernetes, models.NewSmellKubernetes(statefulSet.StatefulSet, statefulSet.StatefulSet.GetObjectKind(), &container, statefulSet.WorkloadPosition, constants.K8S_SEC_RUNASUSER_UNSET))
		}
	}
}

func (statefulSet *StatefulSet) SmellyResourceAndLimit() {
	for _, container := range statefulSet.StatefulSet.Spec.Template.Spec.Containers {
		if container.Resources.Requests == nil {
			statefulSet.SmellKubernetes = append(statefulSet.SmellKubernetes, models.NewSmellKubernetes(statefulSet.StatefulSet, statefulSet.StatefulSet.GetObjectKind(), &container, statefulSet.WorkloadPosition, constants.K8S_SEC_RESREQUESTS_UNSET))
		}
		if container.Resources.Limits == nil {
			statefulSet.SmellKubernetes = append(statefulSet.SmellKubernetes, models.NewSmellKubernetes(statefulSet.StatefulSet, statefulSet.StatefulSet.GetObjectKind(), &container, statefulSet.WorkloadPosition, constants.K8S_SEC_RESLIMITS_UNSET))
		}
	}
}

func (statefulSet *StatefulSet) SmellySecurityContextPrivileged() {
	for _, container := range statefulSet.StatefulSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Privileged == nil {
			statefulSet.SmellKubernetes = append(statefulSet.SmellKubernetes, models.NewSmellKubernetes(statefulSet.StatefulSet, statefulSet.StatefulSet.GetObjectKind(), &container, statefulSet.WorkloadPosition, constants.K8S_SEC_PRIVILEGED_UNSET))
		}
	}
}

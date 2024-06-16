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
			statefulSet.SmellKubernetes = append(statefulSet.SmellKubernetes, models.NewSmellKubernetes(statefulSet.StatefulSet, statefulSet.StatefulSet.GetObjectKind(), &container, statefulSet.WorkloadPosition, constants.K8S_SEC_ROROOTFS))
		}
	}
}

func (statefulSet *StatefulSet) SmellySecurityContextAllowPrivilegeEscalation() {
	for _, container := range statefulSet.StatefulSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			statefulSet.SmellKubernetes = append(statefulSet.SmellKubernetes, models.NewSmellKubernetes(statefulSet.StatefulSet, statefulSet.StatefulSet.GetObjectKind(), &container, statefulSet.WorkloadPosition, constants.K8S_SEC_PRIVESCALATION))
		}
	}
}

func (statefulSet *StatefulSet) SmellySecurityContextCapabilities() {
	for _, container := range statefulSet.StatefulSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			statefulSet.SmellKubernetes = append(statefulSet.SmellKubernetes, models.NewSmellKubernetes(statefulSet.StatefulSet, statefulSet.StatefulSet.GetObjectKind(), &container, statefulSet.WorkloadPosition, constants.K8S_SEC_CAPABILITIES))
		}
	}
}

func (statefulSet *StatefulSet) SmellySecurityContextRunAsUser() {
	for _, container := range statefulSet.StatefulSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			statefulSet.SmellKubernetes = append(statefulSet.SmellKubernetes, models.NewSmellKubernetes(statefulSet.StatefulSet, statefulSet.StatefulSet.GetObjectKind(), &container, statefulSet.WorkloadPosition, constants.K8S_SEC_RUNASUSER))
		}
	}
}

func (statefulSet *StatefulSet) SmellyResourceAndLimit() {
	for _, container := range statefulSet.StatefulSet.Spec.Template.Spec.Containers {
		if container.Resources.Requests == nil {
			statefulSet.SmellKubernetes = append(statefulSet.SmellKubernetes, models.NewSmellKubernetes(statefulSet.StatefulSet, statefulSet.StatefulSet.GetObjectKind(), &container, statefulSet.WorkloadPosition, constants.K8S_SEC_RESREQUESTS))
		}
		if container.Resources.Limits == nil {
			statefulSet.SmellKubernetes = append(statefulSet.SmellKubernetes, models.NewSmellKubernetes(statefulSet.StatefulSet, statefulSet.StatefulSet.GetObjectKind(), &container, statefulSet.WorkloadPosition, constants.K8S_SEC_RESLIMITS))
		}
	}
}

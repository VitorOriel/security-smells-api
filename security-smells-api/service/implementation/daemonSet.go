package implementation

import (
	"security-smells-api/constants"
	"security-smells-api/models"

	appsv1 "k8s.io/api/apps/v1"
)

type DaemonSet struct {
	DaemonSet        *appsv1.DaemonSet
	WorkloadPosition int
	SmellKubernetes  []*models.SmellKubernetes
}

func (daemonSet *DaemonSet) SmellySecurityContextReadOnlyRootFilesystem() {
	for _, container := range daemonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.ReadOnlyRootFilesystem == nil {
			daemonSet.SmellKubernetes = append(daemonSet.SmellKubernetes, models.NewSmellKubernetes(daemonSet.DaemonSet, daemonSet.DaemonSet.GetObjectKind(), &container, daemonSet.WorkloadPosition, constants.K8S_SEC_ROROOTFS_UNSET))
		}
	}
}

func (daemonSet *DaemonSet) SmellySecurityContextAllowPrivilegeEscalation() {
	for _, container := range daemonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			daemonSet.SmellKubernetes = append(daemonSet.SmellKubernetes, models.NewSmellKubernetes(daemonSet.DaemonSet, daemonSet.DaemonSet.GetObjectKind(), &container, daemonSet.WorkloadPosition, constants.K8S_SEC_PRIVESCALATION_UNSET))
		}
	}
}

func (daemonSet *DaemonSet) SmellySecurityContextCapabilities() {
	for _, container := range daemonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			daemonSet.SmellKubernetes = append(daemonSet.SmellKubernetes, models.NewSmellKubernetes(daemonSet.DaemonSet, daemonSet.DaemonSet.GetObjectKind(), &container, daemonSet.WorkloadPosition, constants.K8S_SEC_CAPABILITIES_UNSET))
		}
	}
}

func (daemonSet *DaemonSet) SmellySecurityContextPrivileged() {
	for _, container := range daemonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Privileged == nil {
			daemonSet.SmellKubernetes = append(daemonSet.SmellKubernetes, models.NewSmellKubernetes(daemonSet.DaemonSet, daemonSet.DaemonSet.GetObjectKind(), &container, daemonSet.WorkloadPosition, constants.K8S_SEC_PRIVILEGED_UNSET))
		}
	}
}

func (daemonSet *DaemonSet) SmellySecurityContextRunAsUser() {
	for _, container := range daemonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			daemonSet.SmellKubernetes = append(daemonSet.SmellKubernetes, models.NewSmellKubernetes(daemonSet.DaemonSet, daemonSet.DaemonSet.GetObjectKind(), &container, daemonSet.WorkloadPosition, constants.K8S_SEC_RUNASUSER_UNSET))
		}
	}
}

func (daemonSet *DaemonSet) SmellyResourceAndLimit() {
	for _, container := range daemonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.Resources.Requests == nil || container.Resources.Limits == nil {
			daemonSet.SmellKubernetes = append(daemonSet.SmellKubernetes, models.NewSmellKubernetes(daemonSet.DaemonSet, daemonSet.DaemonSet.GetObjectKind(), &container, daemonSet.WorkloadPosition, constants.K8S_SEC_RESREQUESTS_UNSET))
		}
		if container.Resources.Limits == nil {
			daemonSet.SmellKubernetes = append(daemonSet.SmellKubernetes, models.NewSmellKubernetes(daemonSet.DaemonSet, daemonSet.DaemonSet.GetObjectKind(), &container, daemonSet.WorkloadPosition, constants.K8S_SEC_RESLIMITS_UNSET))
		}
	}
}

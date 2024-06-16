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
			daemonSet.SmellKubernetes = append(daemonSet.SmellKubernetes, models.NewSmellKubernetes(daemonSet.DaemonSet, daemonSet.DaemonSet.GetObjectKind(), &container, daemonSet.WorkloadPosition, constants.K8S_SEC_ROROOTFS))
		}
	}
}

func (daemonSet *DaemonSet) SmellySecurityContextAllowPrivilegeEscalation() {
	for _, container := range daemonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			daemonSet.SmellKubernetes = append(daemonSet.SmellKubernetes, models.NewSmellKubernetes(daemonSet.DaemonSet, daemonSet.DaemonSet.GetObjectKind(), &container, daemonSet.WorkloadPosition, constants.K8S_SEC_PRIVESCALATION))
		}
	}
}

func (daemonSet *DaemonSet) SmellySecurityContextCapabilities() {
	for _, container := range daemonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			daemonSet.SmellKubernetes = append(daemonSet.SmellKubernetes, models.NewSmellKubernetes(daemonSet.DaemonSet, daemonSet.DaemonSet.GetObjectKind(), &container, daemonSet.WorkloadPosition, constants.K8S_SEC_CAPABILITIES))
		}
	}
}

func (daemonSet *DaemonSet) SmellySecurityContextPrivileged() {
	for _, container := range daemonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Privileged == nil {
			daemonSet.SmellKubernetes = append(daemonSet.SmellKubernetes, models.NewSmellKubernetes(daemonSet.DaemonSet, daemonSet.DaemonSet.GetObjectKind(), &container, daemonSet.WorkloadPosition, constants.K8S_SEC_PRIVILEGED))
		}
	}
}

func (daemonSet *DaemonSet) SmellySecurityContextRunAsUser() {
	for _, container := range daemonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			daemonSet.SmellKubernetes = append(daemonSet.SmellKubernetes, models.NewSmellKubernetes(daemonSet.DaemonSet, daemonSet.DaemonSet.GetObjectKind(), &container, daemonSet.WorkloadPosition, constants.K8S_SEC_RUNASUSER))
		}
	}
}

func (daemonSet *DaemonSet) SmellyResourceAndLimit() {
	for _, container := range daemonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.Resources.Requests == nil || container.Resources.Limits == nil {
			daemonSet.SmellKubernetes = append(daemonSet.SmellKubernetes, models.NewSmellKubernetes(daemonSet.DaemonSet, daemonSet.DaemonSet.GetObjectKind(), &container, daemonSet.WorkloadPosition, constants.K8S_SEC_RESREQUESTS))
		}
		if container.Resources.Limits == nil {
			daemonSet.SmellKubernetes = append(daemonSet.SmellKubernetes, models.NewSmellKubernetes(daemonSet.DaemonSet, daemonSet.DaemonSet.GetObjectKind(), &container, daemonSet.WorkloadPosition, constants.K8S_SEC_RESLIMITS))
		}
	}
}

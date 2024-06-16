package implementation

import (
	"security-smells-api/constants"
	"security-smells-api/models"
	"security-smells-api/service/interfaces"

	corev1 "k8s.io/api/core/v1"
)

type Pod struct {
	interfaces.SmellyDeployment
	Pod              *corev1.Pod
	WorkloadPosition int
	SmellKubernetes  []*models.SmellKubernetes
}

func (pod *Pod) SmellySecurityContextAllowPrivilegeEscalation() {
	for _, container := range pod.Pod.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			pod.SmellKubernetes = append(pod.SmellKubernetes, models.NewSmellKubernetes(pod.Pod, pod.Pod.GetObjectKind(), &container, pod.WorkloadPosition, constants.K8S_SEC_PRIVESCALATION))
		}
	}
}

func (pod *Pod) SmellySecurityContextCapabilities() {
	for _, container := range pod.Pod.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			pod.SmellKubernetes = append(pod.SmellKubernetes, models.NewSmellKubernetes(pod.Pod, pod.Pod.GetObjectKind(), &container, pod.WorkloadPosition, constants.K8S_SEC_CAPABILITIES))
		}
	}
}

func (pod *Pod) SmellySecurityContextRunAsUser() {
	for _, container := range pod.Pod.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			pod.SmellKubernetes = append(pod.SmellKubernetes, models.NewSmellKubernetes(pod.Pod, pod.Pod.GetObjectKind(), &container, pod.WorkloadPosition, constants.K8S_SEC_RUNASUSER))
		}
	}
}

func (pod *Pod) SmellyResourceAndLimit() {
	for _, container := range pod.Pod.Spec.Containers {
		if container.Resources.Requests == nil {
			pod.SmellKubernetes = append(pod.SmellKubernetes, models.NewSmellKubernetes(pod.Pod, pod.Pod.GetObjectKind(), &container, pod.WorkloadPosition, constants.K8S_SEC_RESREQUESTS))
		}
		if container.Resources.Limits == nil {
			pod.SmellKubernetes = append(pod.SmellKubernetes, models.NewSmellKubernetes(pod.Pod, pod.Pod.GetObjectKind(), &container, pod.WorkloadPosition, constants.K8S_SEC_RESLIMITS))
		}
	}
}

func (pod *Pod) SmellySecurityContextReadOnlyRootFilesystem() {
	for _, container := range pod.Pod.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.ReadOnlyRootFilesystem == nil {
			pod.SmellKubernetes = append(pod.SmellKubernetes, models.NewSmellKubernetes(pod.Pod, pod.Pod.GetObjectKind(), &container, pod.WorkloadPosition, constants.K8S_SEC_ROROOTFS))
		}
	}
}

func (pod *Pod) SmellySecurityContextPrivileged() {
	for _, container := range pod.Pod.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Privileged == nil {
			pod.SmellKubernetes = append(pod.SmellKubernetes, models.NewSmellKubernetes(pod.Pod, pod.Pod.GetObjectKind(), &container, pod.WorkloadPosition, constants.K8S_SEC_PRIVILEGED))
		}
	}
}

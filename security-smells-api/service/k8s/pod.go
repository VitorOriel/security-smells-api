package k8s

import (
	"security-smells-api/constants"
	"security-smells-api/models"

	corev1 "k8s.io/api/core/v1"
)

type Pod struct {
	Pod              *corev1.Pod
	WorkloadPosition int
	KubernetesSmell  []*models.KubernetesSmell
}

func (pod *Pod) SmellySecurityContextAllowPrivilegeEscalation() {
	for _, container := range pod.Pod.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			pod.KubernetesSmell = append(pod.KubernetesSmell, models.NewKubernetesSmell(pod.Pod, pod.Pod.GetObjectKind(), &container, pod.WorkloadPosition, constants.K8S_SEC_PRIVESCALATION_UNSET))
		}
	}
}

func (pod *Pod) SmellySecurityContextCapabilities() {
	for _, container := range pod.Pod.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			pod.KubernetesSmell = append(pod.KubernetesSmell, models.NewKubernetesSmell(pod.Pod, pod.Pod.GetObjectKind(), &container, pod.WorkloadPosition, constants.K8S_SEC_CAPABILITIES_UNSET))
		}
	}
}

func (pod *Pod) SmellySecurityContextRunAsUser() {
	for _, container := range pod.Pod.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			pod.KubernetesSmell = append(pod.KubernetesSmell, models.NewKubernetesSmell(pod.Pod, pod.Pod.GetObjectKind(), &container, pod.WorkloadPosition, constants.K8S_SEC_RUNASUSER_UNSET))
		}
	}
}

func (pod *Pod) SmellyResourceAndLimit() {
	for _, container := range pod.Pod.Spec.Containers {
		if container.Resources.Requests == nil {
			pod.KubernetesSmell = append(pod.KubernetesSmell, models.NewKubernetesSmell(pod.Pod, pod.Pod.GetObjectKind(), &container, pod.WorkloadPosition, constants.K8S_SEC_RESREQUESTS_UNSET))
		}
		if container.Resources.Limits == nil {
			pod.KubernetesSmell = append(pod.KubernetesSmell, models.NewKubernetesSmell(pod.Pod, pod.Pod.GetObjectKind(), &container, pod.WorkloadPosition, constants.K8S_SEC_RESLIMITS_UNSET))
		}
	}
}

func (pod *Pod) SmellySecurityContextReadOnlyRootFilesystem() {
	for _, container := range pod.Pod.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.ReadOnlyRootFilesystem == nil {
			pod.KubernetesSmell = append(pod.KubernetesSmell, models.NewKubernetesSmell(pod.Pod, pod.Pod.GetObjectKind(), &container, pod.WorkloadPosition, constants.K8S_SEC_ROROOTFS_UNSET))
		}
	}
}

func (pod *Pod) SmellySecurityContextPrivileged() {
	for _, container := range pod.Pod.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Privileged == nil {
			pod.KubernetesSmell = append(pod.KubernetesSmell, models.NewKubernetesSmell(pod.Pod, pod.Pod.GetObjectKind(), &container, pod.WorkloadPosition, constants.K8S_SEC_PRIVILEGED_UNSET))
		}
	}
}

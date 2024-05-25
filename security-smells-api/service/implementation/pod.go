package implementation

import (
	"security-smells-api/models"
	"security-smells-api/service/interfaces"

	corev1 "k8s.io/api/core/v1"
)

type Pod struct {
	interfaces.SmellyDeployment
	Pod              *corev1.Pod
	WorkloadPosition int
	SmellKubernetes  []models.SmellKubernetes
}

func (pod *Pod) SmellySecurityContextAllowPrivilegeEscalation() {
	nameSpace := pod.Pod.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	podName := pod.Pod.GetName()
	kind := pod.Pod.GroupVersionKind().Kind
	for _, container := range pod.Pod.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			smellPod := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: podName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "AllowPrivilegeEscalation not set into " + container.Name + " your container is running with AllowPrivilegeEscalation",
				Suggestion:        "Please add AllowPrivilegeEscalation into " + container.Name + " to avoid running with AllowPrivilegeEscalation",
			}
			pod.SmellKubernetes = append(pod.SmellKubernetes, smellPod)
		}
	}
}

func (pod *Pod) SmellySecurityContextCapabilities() {
	nameSpace := pod.Pod.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	podName := pod.Pod.GetName()
	kind := pod.Pod.GroupVersionKind().Kind
	for _, container := range pod.Pod.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			smellPod := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: podName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "Capabilities not set into " + container.Name + " your container is running without Capabilities",
				Suggestion:        "Please add Capabilities into " + container.Name + " to avoid running without Capabilities",
			}
			pod.SmellKubernetes = append(pod.SmellKubernetes, smellPod)
		}
	}
}

func (pod *Pod) SmellySecurityContextRunAsUser() {
	nameSpace := pod.Pod.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	podName := pod.Pod.GetName()
	kind := pod.Pod.GroupVersionKind().Kind
	for _, container := range pod.Pod.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			smellPod := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: podName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "RunAsUser not set into " + container.Name + " your container is running with root user",
				Suggestion:        "Please add RunAsUser into " + container.Name + " to avoid running with root user",
			}
			pod.SmellKubernetes = append(pod.SmellKubernetes, smellPod)
		}
	}
}

func (Pod *Pod) SmellyResourceAndLimit() {
	nameSpace := Pod.Pod.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	PodName := Pod.Pod.GetName()
	kind := Pod.Pod.GroupVersionKind().Kind
	for _, container := range Pod.Pod.Spec.Containers {
		if container.Resources.Requests == nil {
			smellPod := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: PodName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "Resources not set into " + container.Name + " your container is running without Resources",
				Suggestion:        "Please add Resources into " + container.Name + " to avoid running without Resources",
			}
			Pod.SmellKubernetes = append(Pod.SmellKubernetes, smellPod)
		}
		if container.Resources.Limits == nil {
			smellPod := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: PodName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "Limits not set into " + container.Name + " your container is running without Limits",
				Suggestion:        "Please add Limits into " + container.Name + " to avoid running without Limits",
			}
			Pod.SmellKubernetes = append(Pod.SmellKubernetes, smellPod)
		}
	}
}

func (Pod *Pod) SmellySecurityContextReadOnlyRootFilesystem() {
	nameSpace := Pod.Pod.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	PodName := Pod.Pod.GetName()
	kind := Pod.Pod.GroupVersionKind().Kind
	for _, container := range Pod.Pod.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.ReadOnlyRootFilesystem == nil {
			smellPod := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: PodName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "ReadOnlyRootFilesystem not set into " + container.Name + " your container is running with ReadWriteRootFilesystem",
				Suggestion:        "Please add ReadOnlyRootFilesystem into " + container.Name + " to avoid running with ReadWriteRootFilesystem",
			}
			Pod.SmellKubernetes = append(Pod.SmellKubernetes, smellPod)
		}
	}
}

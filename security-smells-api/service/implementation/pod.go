package implementation

import (
	corev1 "k8s.io/api/core/v1"
	"security-smells-api/models"
	"security-smells-api/service/interfaces"
)

type Pod struct {
	interfaces.SmellyDeployment
	Pod      *corev1.Pod
	SmellPod []models.SmellPod
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
			smellPod := models.SmellPod{
				NameSpace:      nameSpace,
				PodName:        podName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "AllowPrivilegeEscalation not set into " + container.Name + " your container is running with AllowPrivilegeEscalation",
				Suggestion:     "Please add AllowPrivilegeEscalation into " + container.Name + " to avoid running with AllowPrivilegeEscalation",
			}
			pod.SmellPod = append(pod.SmellPod, smellPod)
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
			smellPod := models.SmellPod{
				NameSpace:      nameSpace,
				PodName:        podName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "Capabilities not set into " + container.Name + " your container is running without Capabilities",
				Suggestion:     "Please add Capabilities into " + container.Name + " to avoid running without Capabilities",
			}
			pod.SmellPod = append(pod.SmellPod, smellPod)
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
			smellPod := models.SmellPod{
				NameSpace:      nameSpace,
				PodName:        podName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "RunAsUser not set into " + container.Name + " your container is running with root user",
				Suggestion:     "Please add RunAsUser into " + container.Name + " to avoid running with root user",
			}
			pod.SmellPod = append(pod.SmellPod, smellPod)
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
			smellPod := models.SmellPod{
				NameSpace:      nameSpace,
				PodName:        PodName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "Resources not set into " + container.Name + " your container is running without Resources",
				Suggestion:     "Please add Resources into " + container.Name + " to avoid running without Resources",
			}
			Pod.SmellPod = append(Pod.SmellPod, smellPod)
		}
		if container.Resources.Limits == nil {
			smellPod := models.SmellPod{
				NameSpace:      nameSpace,
				PodName:        PodName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "Limits not set into " + container.Name + " your container is running without Limits",
				Suggestion:     "Please add Limits into " + container.Name + " to avoid running without Limits",
			}
			Pod.SmellPod = append(Pod.SmellPod, smellPod)
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
			smellPod := models.SmellPod{
				NameSpace:      nameSpace,
				PodName:        PodName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "ReadOnlyRootFilesystem not set into " + container.Name + " your container is running with ReadWriteRootFilesystem",
				Suggestion:     "Please add ReadOnlyRootFilesystem into " + container.Name + " to avoid running with ReadWriteRootFilesystem",
			}
			Pod.SmellPod = append(Pod.SmellPod, smellPod)
		}
	}
}

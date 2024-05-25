package implementation

import (
	"security-smells-api/models"

	appsv1 "k8s.io/api/apps/v1"
)

type DaemonSet struct {
	DaemonSet      *appsv1.DaemonSet
	SmellDaemonSet []models.SmellDaemonSet
}

func (daemonSet *DaemonSet) SmellySecurityContextReadOnlyRootFilesystem() {
	nameSpace := daemonSet.DaemonSet.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	daemonSetName := daemonSet.DaemonSet.GetName()
	kind := daemonSet.DaemonSet.GroupVersionKind().Kind
	for _, container := range daemonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.ReadOnlyRootFilesystem == nil {
			smellDemonSet := models.SmellDaemonSet{
				NameSpace:      nameSpace,
				DemonSetName:   daemonSetName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "ReadOnlyRootFilesystem not set into " + container.Name + " your container is running with ReadWriteRootFilesystem",
				Suggestion:     "Please add ReadOnlyRootFilesystem into " + container.Name + " to avoid running with ReadWriteRootFilesystem",
			}
			daemonSet.SmellDaemonSet = append(daemonSet.SmellDaemonSet, smellDemonSet)
		}
	}
}

func (daemonSet *DaemonSet) SmellySecurityContextAllowPrivilegeEscalation() {
	nameSpace := daemonSet.DaemonSet.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	daemonSetName := daemonSet.DaemonSet.GetName()
	kind := daemonSet.DaemonSet.GroupVersionKind().Kind
	for _, container := range daemonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			smellDemonSet := models.SmellDaemonSet{
				NameSpace:      nameSpace,
				DemonSetName:   daemonSetName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "AllowPrivilegeEscalation not set into " + container.Name + " your container is running with AllowPrivilegeEscalation",
				Suggestion:     "Please add AllowPrivilegeEscalation into " + container.Name + " to avoid running with AllowPrivilegeEscalation",
			}
			daemonSet.SmellDaemonSet = append(daemonSet.SmellDaemonSet, smellDemonSet)
		}
	}
}

func (daemonSet *DaemonSet) SmellySecurityContextCapabilities() {
	nameSpace := daemonSet.DaemonSet.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	daemonSetName := daemonSet.DaemonSet.GetName()
	kind := daemonSet.DaemonSet.GroupVersionKind().Kind
	for _, container := range daemonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			smellDemonSet := models.SmellDaemonSet{
				NameSpace:      nameSpace,
				DemonSetName:   daemonSetName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "Capabilities not set into " + container.Name + " your container is running with all capabilities",
				Suggestion:     "Please add Capabilities into " + container.Name + " to avoid running with all capabilities",
			}
			daemonSet.SmellDaemonSet = append(daemonSet.SmellDaemonSet, smellDemonSet)
		}
	}
}

func (daemonSet *DaemonSet) SmellySecurityContextPrivileged() {
	nameSpace := daemonSet.DaemonSet.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	daemonSetName := daemonSet.DaemonSet.GetName()
	kind := daemonSet.DaemonSet.GroupVersionKind().Kind
	for _, container := range daemonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Privileged == nil {
			smellDemonSet := models.SmellDaemonSet{
				NameSpace:      nameSpace,
				DemonSetName:   daemonSetName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "Privileged not set into " + container.Name + " your container is running with Privileged",
				Suggestion:     "Please add Privileged into " + container.Name + " to avoid running with Privileged",
			}
			daemonSet.SmellDaemonSet = append(daemonSet.SmellDaemonSet, smellDemonSet)
		}
	}
}

func (daemonSet *DaemonSet) SmellySecurityContextRunAsUser() {
	nameSpace := daemonSet.DaemonSet.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	daemonSetName := daemonSet.DaemonSet.GetName()
	kind := daemonSet.DaemonSet.GroupVersionKind().Kind
	for _, container := range daemonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			smellDemonSet := models.SmellDaemonSet{
				NameSpace:      nameSpace,
				DemonSetName:   daemonSetName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "RunAsUser not set into " + container.Name + " your container is running with root user",
				Suggestion:     "Please add RunAsUser into " + container.Name + " to avoid running with root user",
			}
			daemonSet.SmellDaemonSet = append(daemonSet.SmellDaemonSet, smellDemonSet)
		}
	}
}

func (daemonSet *DaemonSet) SmellyResourceAndLimit() {
	nameSpace := daemonSet.DaemonSet.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	daemonSetName := daemonSet.DaemonSet.GetName()
	kind := daemonSet.DaemonSet.GroupVersionKind().Kind
	for _, container := range daemonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.Resources.Requests == nil || container.Resources.Limits == nil {
			smellDemonSet := models.SmellDaemonSet{
				NameSpace:      nameSpace,
				DemonSetName:   daemonSetName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "Resources and Limits not set into " + container.Name + " your container is running without resource and limit",
				Suggestion:     "Please add Resources and Limits into " + container.Name + " to avoid running without resource and limit",
			}
			daemonSet.SmellDaemonSet = append(daemonSet.SmellDaemonSet, smellDemonSet)
		}
	}
}

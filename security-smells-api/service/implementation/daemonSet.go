package implementation

import (
	appsv1 "k8s.io/api/apps/v1"
	"security-smells-api/models"
)

type DaemonSet struct {
	DaemonSet      *appsv1.DaemonSet
	SmellDaemonSet []models.SmellDaemonSet
}

func (demonSet *DaemonSet) SmellySecurityContextReadOnlyRootFilesystem() {
	nameSpace := demonSet.DaemonSet.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	demonSetName := demonSet.DaemonSet.GetName()
	kind := demonSet.DaemonSet.GroupVersionKind().Kind
	for _, container := range demonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.ReadOnlyRootFilesystem == nil {
			smellDemonSet := models.SmellDaemonSet{
				NameSpace:      nameSpace,
				DemonSetName:   demonSetName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "ReadOnlyRootFilesystem not set into " + container.Name + " your container is running with ReadWriteRootFilesystem",
				Suggestion:     "Please add ReadOnlyRootFilesystem into " + container.Name + " to avoid running with ReadWriteRootFilesystem",
			}
			demonSet.SmellDaemonSet = append(demonSet.SmellDaemonSet, smellDemonSet)
		}
	}
}

func (demonSet *DaemonSet) SmellySecurityContextAllowPrivilegeEscalation() {
	nameSpace := demonSet.DaemonSet.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	demonSetName := demonSet.DaemonSet.GetName()
	kind := demonSet.DaemonSet.GroupVersionKind().Kind
	for _, container := range demonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			smellDemonSet := models.SmellDaemonSet{
				NameSpace:      nameSpace,
				DemonSetName:   demonSetName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "AllowPrivilegeEscalation not set into " + container.Name + " your container is running with AllowPrivilegeEscalation",
				Suggestion:     "Please add AllowPrivilegeEscalation into " + container.Name + " to avoid running with AllowPrivilegeEscalation",
			}
			demonSet.SmellDaemonSet = append(demonSet.SmellDaemonSet, smellDemonSet)
		}
	}
}

func (demonSet *DaemonSet) SmellySecurityContextCapabilities() {
	nameSpace := demonSet.DaemonSet.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	demonSetName := demonSet.DaemonSet.GetName()
	kind := demonSet.DaemonSet.GroupVersionKind().Kind
	for _, container := range demonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			smellDemonSet := models.SmellDaemonSet{
				NameSpace:      nameSpace,
				DemonSetName:   demonSetName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "Capabilities not set into " + container.Name + " your container is running with all capabilities",
				Suggestion:     "Please add Capabilities into " + container.Name + " to avoid running with all capabilities",
			}
			demonSet.SmellDaemonSet = append(demonSet.SmellDaemonSet, smellDemonSet)
		}
	}
}

func (demonSet *DaemonSet) SmellySecurityContextPrivileged() {
	nameSpace := demonSet.DaemonSet.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	demonSetName := demonSet.DaemonSet.GetName()
	kind := demonSet.DaemonSet.GroupVersionKind().Kind
	for _, container := range demonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Privileged == nil {
			smellDemonSet := models.SmellDaemonSet{
				NameSpace:      nameSpace,
				DemonSetName:   demonSetName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "Privileged not set into " + container.Name + " your container is running with Privileged",
				Suggestion:     "Please add Privileged into " + container.Name + " to avoid running with Privileged",
			}
			demonSet.SmellDaemonSet = append(demonSet.SmellDaemonSet, smellDemonSet)
		}
	}
}

func (demonSet *DaemonSet) SmellySecurityContextRunAsUser() {
	nameSpace := demonSet.DaemonSet.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	demonSetName := demonSet.DaemonSet.GetName()
	kind := demonSet.DaemonSet.GroupVersionKind().Kind
	for _, container := range demonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			smellDemonSet := models.SmellDaemonSet{
				NameSpace:      nameSpace,
				DemonSetName:   demonSetName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "RunAsUser not set into " + container.Name + " your container is running with root user",
				Suggestion:     "Please add RunAsUser into " + container.Name + " to avoid running with root user",
			}
			demonSet.SmellDaemonSet = append(demonSet.SmellDaemonSet, smellDemonSet)
		}
	}
}

func (demonSet *DaemonSet) SmellyResourceAndLimit() {
	nameSpace := demonSet.DaemonSet.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	demonSetName := demonSet.DaemonSet.GetName()
	kind := demonSet.DaemonSet.GroupVersionKind().Kind
	for _, container := range demonSet.DaemonSet.Spec.Template.Spec.Containers {
		if container.Resources.Requests == nil || container.Resources.Limits == nil {
			smellDemonSet := models.SmellDaemonSet{
				NameSpace:      nameSpace,
				DemonSetName:   demonSetName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "Resources and Limits not set into " + container.Name + " your container is running without resource and limit",
				Suggestion:     "Please add Resources and Limits into " + container.Name + " to avoid running without resource and limit",
			}
			demonSet.SmellDaemonSet = append(demonSet.SmellDaemonSet, smellDemonSet)
		}
	}
}

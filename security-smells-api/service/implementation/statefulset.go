package implementation

import (
	"security-smells-api/models"
	"security-smells-api/service/interfaces"

	v1 "k8s.io/api/apps/v1"
)

type StatefulSet struct {
	interfaces.SmellyStatefulSet
	StatefulSet      *v1.StatefulSet
	WorkloadPosition int
	SmellKubernetes  []models.SmellKubernetes
}

func (statefulset *StatefulSet) SmellySecurityContextReadOnlyRootFilesystem() {
	nameSpace := statefulset.StatefulSet.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	statefulsetName := statefulset.StatefulSet.GetName()
	kind := statefulset.StatefulSet.GroupVersionKind().Kind
	for _, container := range statefulset.StatefulSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.ReadOnlyRootFilesystem == nil {
			smellStatefulset := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: statefulsetName,
				WorkloadPosition:  statefulset.WorkloadPosition,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "ReadOnlyRootFilesystem not set into " + container.Name + " your container is running with ReadWriteRootFilesystem",
				Suggestion:        "Please add ReadOnlyRootFilesystem into " + container.Name + " to avoid running with ReadWriteRootFilesystem",
			}
			statefulset.SmellKubernetes = append(statefulset.SmellKubernetes, smellStatefulset)
		}
	}
}

func (statefulset *StatefulSet) SmellySecurityContextAllowPrivilegeEscalation() {
	nameSpace := statefulset.StatefulSet.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	statefulsetName := statefulset.StatefulSet.GetName()
	kind := statefulset.StatefulSet.GroupVersionKind().Kind
	for _, container := range statefulset.StatefulSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			smellStatefulset := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: statefulsetName,
				WorkloadPosition:  statefulset.WorkloadPosition,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "AllowPrivilegeEscalation not set into " + container.Name + " your container is running with AllowPrivilegeEscalation",
				Suggestion:        "Please add AllowPrivilegeEscalation into " + container.Name + " to avoid running with AllowPrivilegeEscalation",
			}
			statefulset.SmellKubernetes = append(statefulset.SmellKubernetes, smellStatefulset)
		}
	}
}

func (statefulset *StatefulSet) SmellySecurityContextCapabilities() {
	nameSpace := statefulset.StatefulSet.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	statefulsetName := statefulset.StatefulSet.GetName()
	kind := statefulset.StatefulSet.GroupVersionKind().Kind
	for _, container := range statefulset.StatefulSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			smellStatefulset := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: statefulsetName,
				WorkloadPosition:  statefulset.WorkloadPosition,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "Capabilities not set into " + container.Name + " your container is running with all the capabilities",
				Suggestion:        "Please add Capabilities into " + container.Name + " to avoid running with all the capabilities",
			}
			statefulset.SmellKubernetes = append(statefulset.SmellKubernetes, smellStatefulset)
		}
	}
}

func (statefulset *StatefulSet) SmellySecurityContextRunAsUser() {
	nameSpace := statefulset.StatefulSet.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	statefulsetName := statefulset.StatefulSet.GetName()
	kind := statefulset.StatefulSet.GroupVersionKind().Kind
	for _, container := range statefulset.StatefulSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			smellStatefulset := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: statefulsetName,
				WorkloadPosition:  statefulset.WorkloadPosition,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "RunAsUser not set into " + container.Name + " your container is running with root user",
				Suggestion:        "Please add RunAsUser into " + container.Name + " to avoid running with root user",
			}
			statefulset.SmellKubernetes = append(statefulset.SmellKubernetes, smellStatefulset)
		}
	}
}

func (statefulset *StatefulSet) SmellyResourceAndLimit() {
	nameSpace := statefulset.StatefulSet.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	statefulsetName := statefulset.StatefulSet.GetName()
	kind := statefulset.StatefulSet.GroupVersionKind().Kind
	for _, container := range statefulset.StatefulSet.Spec.Template.Spec.Containers {
		if container.Resources.Requests == nil || container.Resources.Limits == nil {
			smellStatefulset := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: statefulsetName,
				WorkloadPosition:  statefulset.WorkloadPosition,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "Resource and Limit not set into " + container.Name + " your container is running without resource and limit",
				Suggestion:        "Please add Resource and Limit into " + container.Name + " to avoid running without resource and limit",
			}
			statefulset.SmellKubernetes = append(statefulset.SmellKubernetes, smellStatefulset)
		}
	}
}

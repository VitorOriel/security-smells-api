package implementation

import (
	"security-smells-api/models"
	"security-smells-api/service/interfaces"

	appsv1 "k8s.io/api/apps/v1"
)

type ReplicaSet struct {
	interfaces.SmellyReplicaSet
	ReplicaSet       *appsv1.ReplicaSet
	WorkloadPosition int
	SmellKubernetes  []models.SmellKubernetes
}

func (replicaSet *ReplicaSet) SmellySecurityContextReadOnlyRootFilesystem() {
	nameSpace := replicaSet.ReplicaSet.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	replicaSetName := replicaSet.ReplicaSet.GetName()
	kind := replicaSet.ReplicaSet.GroupVersionKind().Kind
	for _, container := range replicaSet.ReplicaSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.ReadOnlyRootFilesystem == nil {
			smellReplicaSet := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: replicaSetName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "ReadOnlyRootFilesystem not set into " + container.Name + " your container is running with ReadWriteRootFilesystem",
				Suggestion:        "Please add ReadOnlyRootFilesystem into " + container.Name + " to avoid running with ReadWriteRootFilesystem",
			}
			replicaSet.SmellKubernetes = append(replicaSet.SmellKubernetes, smellReplicaSet)
		}
	}
}

func (replicaSet *ReplicaSet) SmellySecurityContextAllowPrivilegeEscalation() {
	nameSpace := replicaSet.ReplicaSet.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	replicaSetName := replicaSet.ReplicaSet.GetName()
	kind := replicaSet.ReplicaSet.GroupVersionKind().Kind
	for _, container := range replicaSet.ReplicaSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			smellReplicaSet := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: replicaSetName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "AllowPrivilegeEscalation not set into " + container.Name + " your container is running with AllowPrivilegeEscalation",
				Suggestion:        "Please add AllowPrivilegeEscalation into " + container.Name + " to avoid running with AllowPrivilegeEscalation",
			}
			replicaSet.SmellKubernetes = append(replicaSet.SmellKubernetes, smellReplicaSet)
		}
	}
}

func (replicaSet *ReplicaSet) SmellySecurityContextCapabilities() {
	nameSpace := replicaSet.ReplicaSet.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	replicaSetName := replicaSet.ReplicaSet.GetName()
	kind := replicaSet.ReplicaSet.GroupVersionKind().Kind
	for _, container := range replicaSet.ReplicaSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			smellReplicaSet := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: replicaSetName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "Capabilities not set into " + container.Name + " your container is running with full capabilities",
				Suggestion:        "Please add capabilities into " + container.Name + " to avoid running with full capabilities",
			}
			replicaSet.SmellKubernetes = append(replicaSet.SmellKubernetes, smellReplicaSet)
		}

	}
}

func (replicaSet *ReplicaSet) SmellySecurityContextRunAsUser() {
	nameSpace := replicaSet.ReplicaSet.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	replicaSetName := replicaSet.ReplicaSet.GetName()
	kind := replicaSet.ReplicaSet.GroupVersionKind().Kind
	for _, container := range replicaSet.ReplicaSet.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			smellReplicaSet := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: replicaSetName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "RunAsUser not set into " + container.Name + " your container is running as root",
				Suggestion:        "Please add runAsUser into " + container.Name + " to avoid running as root",
			}
			replicaSet.SmellKubernetes = append(replicaSet.SmellKubernetes, smellReplicaSet)
		}
	}
}

func (replicaSet *ReplicaSet) SmellyResourceAndLimit() {
	// Check if the replicaSet has resource limits set
	//verify for all containers
	nameSpace := replicaSet.ReplicaSet.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	replicaSetName := replicaSet.ReplicaSet.GetName()
	kind := replicaSet.ReplicaSet.GroupVersionKind().Kind
	for _, container := range replicaSet.ReplicaSet.Spec.Template.Spec.Containers {
		if container.Resources.Limits == nil {
			smellReplicaSet := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: replicaSetName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "Resource limits not set for container " + container.Name,
				Suggestion:        "Set resource limits for container " + container.Name,
			}
			replicaSet.SmellKubernetes = append(replicaSet.SmellKubernetes, smellReplicaSet)
		}
		if container.Resources.Requests == nil {
			smellReplicaSet := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: replicaSetName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "Resource requests not set for container " + container.Name,
				Suggestion:        "Set resource requests for container " + container.Name,
			}
			replicaSet.SmellKubernetes = append(replicaSet.SmellKubernetes, smellReplicaSet)
		}
	}
}

func (replicaSet *ReplicaSet) SmellySecurityContext() {
	//TODO implement me
	panic("implement me")
}

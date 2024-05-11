package implementation

import (
	"security-smells-api/models"
	"security-smells-api/service/interfaces"

	appsv1 "k8s.io/api/apps/v1"
)

type ReplicaSet struct {
	interfaces.SmellyReplicaSet
	ReplicaSet      *appsv1.ReplicaSet
	SmellReplicaSet []models.SmellReplicaSet
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
			smellReplicaSet := models.SmellReplicaSet{
				NameSpace:      nameSpace,
				ReplicaSetName: replicaSetName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "ReadOnlyRootFilesystem not set into " + container.Name + " your container is running with ReadWriteRootFilesystem",
				Suggestion:     "Please add ReadOnlyRootFilesystem into " + container.Name + " to avoid running with ReadWriteRootFilesystem",
			}
			replicaSet.SmellReplicaSet = append(replicaSet.SmellReplicaSet, smellReplicaSet)
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
			smellReplicaSet := models.SmellReplicaSet{
				NameSpace:      nameSpace,
				ReplicaSetName: replicaSetName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "AllowPrivilegeEscalation not set into " + container.Name + " your container is running with AllowPrivilegeEscalation",
				Suggestion:     "Please add AllowPrivilegeEscalation into " + container.Name + " to avoid running with AllowPrivilegeEscalation",
			}
			replicaSet.SmellReplicaSet = append(replicaSet.SmellReplicaSet, smellReplicaSet)
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
			smellReplicaSet := models.SmellReplicaSet{
				NameSpace:      nameSpace,
				ReplicaSetName: replicaSetName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "Capabilities not set into " + container.Name + " your container is running with full capabilities",
				Suggestion:     "Please add capabilities into " + container.Name + " to avoid running with full capabilities",
			}
			replicaSet.SmellReplicaSet = append(replicaSet.SmellReplicaSet, smellReplicaSet)
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
			smellReplicaSet := models.SmellReplicaSet{
				NameSpace:      nameSpace,
				ReplicaSetName: replicaSetName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "RunAsUser not set into " + container.Name + " your container is running as root",
				Suggestion:     "Please add runAsUser into " + container.Name + " to avoid running as root",
			}
			replicaSet.SmellReplicaSet = append(replicaSet.SmellReplicaSet, smellReplicaSet)
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
			smellReplicaSet := models.SmellReplicaSet{
				NameSpace:      nameSpace,
				ReplicaSetName: replicaSetName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "Resource limits not set for container " + container.Name,
				Suggestion:     "Set resource limits for container " + container.Name,
			}
			replicaSet.SmellReplicaSet = append(replicaSet.SmellReplicaSet, smellReplicaSet)
		}
		if container.Resources.Requests == nil {
			smellReplicaSet := models.SmellReplicaSet{
				NameSpace:      nameSpace,
				ReplicaSetName: replicaSetName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "Resource requests not set for container " + container.Name,
				Suggestion:     "Set resource requests for container " + container.Name,
			}
			replicaSet.SmellReplicaSet = append(replicaSet.SmellReplicaSet, smellReplicaSet)
		}
	}
}

func (replicaSet *ReplicaSet) SmellySecurityContext() {
	//TODO implement me
	panic("implement me")
}

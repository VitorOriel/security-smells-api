package models

import (
	"fmt"
	"security-smells-api/constants"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type KubernetesSmell struct {
	Namespace         string                   `json:"namespace" validate:"required"`
	WorkloadKind      string                   `json:"workload_kind" validate:"required"`
	WorkloadLabelName string                   `json:"workload_label_name" validate:"required"`
	WorkloadPosition  int                      `json:"workload_position" validate:"required"`
	Rule              constants.KubernetesRule `json:"rule" validate:"required"`
	Message           string                   `json:"message" validate:"required"`
	Suggestion        string                   `json:"suggestion" validate:"required"`
}

func NewKubernetesSmell(
	k8sObj metav1.Object,
	k8sKind schema.ObjectKind,
	container *corev1.Container,
	workloadPosition int,
	rule constants.KubernetesRule,
) *KubernetesSmell {
	namespace := k8sObj.GetNamespace()
	if namespace == "" {
		namespace = "default"
	}
	kubernetesSmells := &KubernetesSmell{
		Namespace:         namespace,
		WorkloadKind:      k8sKind.GroupVersionKind().Kind,
		WorkloadLabelName: k8sObj.GetName(),
		WorkloadPosition:  workloadPosition,
		Rule:              rule,
	}
	kubernetesSmells.setMessageAndSuggestionByRule(container, rule)
	return kubernetesSmells
}

func (kubernetesSmells *KubernetesSmell) setMessageAndSuggestionByRule(container *corev1.Container, rule constants.KubernetesRule) {
	switch rule {
	case constants.K8S_SEC_PRIVESCALATION_UNSET:
		kubernetesSmells.Message = fmt.Sprintf("AllowPrivilegeEscalation not set into %s your container is running with AllowPrivilegeEscalation", container.Name)
		kubernetesSmells.Suggestion = fmt.Sprintf("Please add AllowPrivilegeEscalation into %s to avoid running with AllowPrivilegeEscalation", container.Name)
	case constants.K8S_SEC_CAPABILITIES_UNSET:
		kubernetesSmells.Message = fmt.Sprintf("Capabilities not set into %s your container is running without Capabilities", container.Name)
		kubernetesSmells.Suggestion = fmt.Sprintf("Please add Capabilities into %s to avoid running without Capabilities", container.Name)
	case constants.K8S_SEC_RUNASUSER_UNSET:
		kubernetesSmells.Message = fmt.Sprintf("RunAsUser not set into %s your container is running with root user", container.Name)
		kubernetesSmells.Suggestion = fmt.Sprintf("Please add RunAsUser into %s to avoid running with root user", container.Name)
	case constants.K8S_SEC_RESREQUESTS_UNSET:
		kubernetesSmells.Message = fmt.Sprintf("Resources.Requests not set into %s your container is running without Resources.Requests", container.Name)
		kubernetesSmells.Suggestion = fmt.Sprintf("Please add Resources.Requests into %s to avoid running without Resources.Requests", container.Name)
	case constants.K8S_SEC_RESLIMITS_UNSET:
		kubernetesSmells.Message = fmt.Sprintf("Resource.Limits not set into %s your container is running without Resource.Limits", container.Name)
		kubernetesSmells.Suggestion = fmt.Sprintf("Please add Resource.Limits into %s to avoid running without Resource.Limits", container.Name)
	case constants.K8S_SEC_ROROOTFS_UNSET:
		kubernetesSmells.Message = fmt.Sprintf("ReadOnlyRootFilesystem not set into %s your container is running with ReadWriteRootFilesystem", container.Name)
		kubernetesSmells.Suggestion = fmt.Sprintf("Please add ReadOnlyRootFilesystem into %s to avoid running with ReadWriteRootFilesystem", container.Name)
	case constants.K8S_SEC_PRIVILEGED_UNSET:
		kubernetesSmells.Message = fmt.Sprintf("Privileged not set into %s your container is running with Privileged", container.Name)
		kubernetesSmells.Suggestion = fmt.Sprintf("Please add Privileged into %s to avoid running with Privileged", container.Name)
	}
}

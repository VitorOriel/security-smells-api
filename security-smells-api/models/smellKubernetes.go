package models

import (
	"fmt"
	"security-smells-api/constants"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type SmellKubernetes struct {
	Namespace         string                   `json:"namespace" validate:"required"`
	WorkloadKind      string                   `json:"workload_kind" validate:"required"`
	WorkloadLabelName string                   `json:"workload_label_name" validate:"required"`
	WorkloadPosition  int                      `json:"workload_position" validate:"required"`
	Rule              constants.KubernetesRule `json:"rule" validate:"required"`
	Message           string                   `json:"message" validate:"required"`
	Suggestion        string                   `json:"suggestion" validate:"required"`
}

func NewSmellKubernetes(
	k8sObj metav1.Object,
	k8sKind schema.ObjectKind,
	container *corev1.Container,
	workloadPosition int,
	rule constants.KubernetesRule,
) *SmellKubernetes {
	namespace := k8sObj.GetNamespace()
	if namespace == "" {
		namespace = "default"
	}
	smellKubernetes := &SmellKubernetes{
		Namespace:         namespace,
		WorkloadKind:      k8sKind.GroupVersionKind().Kind,
		WorkloadLabelName: k8sObj.GetName(),
		Rule:              rule,
	}
	smellKubernetes.setMessageAndSuggestionByRule(container, rule)
	return smellKubernetes
}

func (smellKubernetes *SmellKubernetes) setMessageAndSuggestionByRule(container *corev1.Container, rule constants.KubernetesRule) {
	switch rule {
	case constants.K8S_SEC_PRIVESCALATION:
		smellKubernetes.Message = fmt.Sprintf("AllowPrivilegeEscalation not set into %s your container is running with AllowPrivilegeEscalation", container.Name)
		smellKubernetes.Suggestion = fmt.Sprintf("Please add AllowPrivilegeEscalation into %s to avoid running with AllowPrivilegeEscalation", container.Name)
	case constants.K8S_SEC_CAPABILITIES:
		smellKubernetes.Message = fmt.Sprintf("Capabilities not set into %s your container is running without Capabilities", container.Name)
		smellKubernetes.Suggestion = fmt.Sprintf("Please add Capabilities into %s to avoid running without Capabilities", container.Name)
	case constants.K8S_SEC_RUNASUSER:
		smellKubernetes.Message = fmt.Sprintf("RunAsUser not set into %s your container is running with root user", container.Name)
		smellKubernetes.Suggestion = fmt.Sprintf("Please add RunAsUser into %s to avoid running with root user", container.Name)
	case constants.K8S_SEC_RESREQUESTS:
		smellKubernetes.Message = fmt.Sprintf("Resources.Requests not set into %s your container is running without Resources.Requests", container.Name)
		smellKubernetes.Suggestion = fmt.Sprintf("Please add Resources.Requests into %s to avoid running without Resources.Requests", container.Name)
	case constants.K8S_SEC_RESLIMITS:
		smellKubernetes.Message = fmt.Sprintf("Resource.Limits not set into %s your container is running without Resource.Limits", container.Name)
		smellKubernetes.Suggestion = fmt.Sprintf("Please add Resource.Limits into %s to avoid running without Resource.Limits", container.Name)
	case constants.K8S_SEC_ROROOTFS:
		smellKubernetes.Message = fmt.Sprintf("ReadOnlyRootFilesystem not set into %s your container is running with ReadWriteRootFilesystem", container.Name)
		smellKubernetes.Suggestion = fmt.Sprintf("Please add ReadOnlyRootFilesystem into %s to avoid running with ReadWriteRootFilesystem", container.Name)
	case constants.K8S_SEC_PRIVILEGED:
		smellKubernetes.Message = fmt.Sprintf("Privileged not set into %s your container is running with Privileged", container.Name)
		smellKubernetes.Suggestion = fmt.Sprintf("Please add Privileged into %s to avoid running with Privileged", container.Name)
	}
}

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
	kubernetesSmell := &KubernetesSmell{
		Namespace:         namespace,
		WorkloadKind:      k8sKind.GroupVersionKind().Kind,
		WorkloadLabelName: k8sObj.GetName(),
		WorkloadPosition:  workloadPosition,
		Rule:              rule,
	}
	kubernetesSmell.setMessageAndSuggestionByRule(container, rule)
	return kubernetesSmell
}

func (kubernetesSmell *KubernetesSmell) setMessageAndSuggestionByRule(container *corev1.Container, rule constants.KubernetesRule) {
	switch rule {
	case constants.K8S_SEC_PRIVESCALATION_UNSET:
		kubernetesSmell.Message = fmt.Sprintf("AllowPrivilegeEscalation is not set in container %s. Your container may run with AllowPrivilegeEscalation enabled.", container.Name)
		kubernetesSmell.Suggestion = fmt.Sprintf("Please add AllowPrivilegeEscalation into %s to avoid running with AllowPrivilegeEscalation", container.Name)
	case constants.K8S_SEC_CAPABILITIES_UNSET:
		kubernetesSmell.Message = fmt.Sprintf("Capabilities not set into %s your container may run with unnecessary capabilities", container.Name)
		kubernetesSmell.Suggestion = fmt.Sprintf("Please add Capabilities drop ALL into %s to avoid running with unnecessary capabilities", container.Name)
	case constants.K8S_SEC_RUNASUSER_UNSET:
		kubernetesSmell.Message = fmt.Sprintf("RunAsUser not set into %s your container may run with root user", container.Name)
		kubernetesSmell.Suggestion = fmt.Sprintf("Please add RunAsUser into %s to specify which user the container should run as", container.Name)
	case constants.K8S_SEC_RUNASNONROOT_UNSET:
		kubernetesSmell.Message = fmt.Sprintf("RunAsNonRoot not set into %s your container is able to run with root user", container.Name)
		kubernetesSmell.Suggestion = fmt.Sprintf("Please add RunAsNonRoot into %s to enforce not running with root user", container.Name)
	case constants.K8S_SEC_RESREQUESTS_UNSET:
		kubernetesSmell.Message = fmt.Sprintf("Resources.Requests not set into %s your container is running without Resources.Requests", container.Name)
		kubernetesSmell.Suggestion = fmt.Sprintf("Please add Resources.Requests into %s to avoid running without Resources.Requests", container.Name)
	case constants.K8S_SEC_RESLIMITS_UNSET:
		kubernetesSmell.Message = fmt.Sprintf("Resource.Limits not set into %s your container is running without Resource.Limits", container.Name)
		kubernetesSmell.Suggestion = fmt.Sprintf("Please add Resource.Limits into %s to avoid running without Resource.Limits", container.Name)
	case constants.K8S_SEC_ROROOTFS_UNSET:
		kubernetesSmell.Message = fmt.Sprintf("ReadOnlyRootFilesystem not set into %s your container is running with ReadWriteRootFilesystem", container.Name)
		kubernetesSmell.Suggestion = fmt.Sprintf("Please add ReadOnlyRootFilesystem into %s to avoid running with ReadWriteRootFilesystem", container.Name)
	case constants.K8S_SEC_PRIVESCALATION_VALUE:
		kubernetesSmell.Message = fmt.Sprintf("AllowPrivilegeEscalation value set to true into %s container", container.Name)
		kubernetesSmell.Suggestion = fmt.Sprintf("You should set the value into %s container to 'AllowPrivilegeEscalation: false'", container.Name)
	case constants.K8S_SEC_CAPABILITIES_VALUE:
		kubernetesSmell.Message = fmt.Sprintf("Capabilities value not set properly into %s container", container.Name)
		kubernetesSmell.Suggestion = fmt.Sprintf("You should set the value into %s container to Capabilities:\n  - \"ALL\"", container.Name)
	case constants.K8S_SEC_RUNASUSER_VALUE:
		kubernetesSmell.Message = fmt.Sprintf("RunAsUser value set to zero into %s container", container.Name)
		kubernetesSmell.Suggestion = fmt.Sprintf("You should set the value into %s container to RunAsUser greater than zero", container.Name)
	case constants.K8S_SEC_RUNASNONROOT_VALUE:
		kubernetesSmell.Message = fmt.Sprintf("RunAsNonRoot value set to false into %s container", container.Name)
		kubernetesSmell.Suggestion = fmt.Sprintf("You should set the value into %s container to 'RunAsNonRoot: true'", container.Name)
	case constants.K8S_SEC_ROROOTFS_VALUE:
		kubernetesSmell.Message = fmt.Sprintf("ReadOnlyRootFilesystem value set to false into %s container", container.Name)
		kubernetesSmell.Suggestion = fmt.Sprintf("You should set the value into %s container to 'ReadOnlyRootFilesystem: true'", container.Name)
	case constants.K8S_SEC_PRIVILEGED_VALUE:
		kubernetesSmell.Message = fmt.Sprintf("Privileged setted to 'true' into %s your container is running with Privileged", container.Name)
		kubernetesSmell.Suggestion = fmt.Sprintf("You should set 'privileged: false' into %s to avoid running with Privileged", container.Name)
	}
}

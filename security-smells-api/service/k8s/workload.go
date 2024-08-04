package k8s

import (
	"security-smells-api/constants"
	"security-smells-api/models"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type K8sWorkload interface {
	GetKubernetesSmells() []*models.KubernetesSmell
	SmellySecurityContextRunAsUser(*corev1.PodSpec)
	SmellySecurityContextCapabilities(*corev1.PodSpec)
	SmellySecurityContextAllowPrivilegeEscalation(*corev1.PodSpec)
	SmellySecurityContextReadOnlyRootFilesystem(*corev1.PodSpec)
	SmellySecurityContextPrivileged(*corev1.PodSpec)
	SmellyResourceAndLimit(*corev1.PodSpec)
}

type k8sWorkload struct {
	object           metav1.Object
	kind             schema.ObjectKind
	workloadPosition int
	kubernetesSmells []*models.KubernetesSmell
}

func NewK8sWorkload(workloadPosition int) K8sWorkload {
	return &k8sWorkload{workloadPosition: workloadPosition}
}

func (w *k8sWorkload) GetKubernetesSmells() []*models.KubernetesSmell {
	return w.kubernetesSmells
}

func (w *k8sWorkload) SmellySecurityContextRunAsUser(spec *corev1.PodSpec) {
	if spec.SecurityContext != nil && spec.SecurityContext.RunAsUser != nil {
		return
	}
	for _, container := range spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			w.kubernetesSmells = append(
				w.kubernetesSmells,
				models.NewKubernetesSmell(w.object, w.kind, &container, w.workloadPosition, constants.K8S_SEC_RUNASUSER_UNSET),
			)
		}
	}
}

func (w *k8sWorkload) SmellySecurityContextCapabilities(spec *corev1.PodSpec) {
	for _, container := range spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			w.kubernetesSmells = append(
				w.kubernetesSmells,
				models.NewKubernetesSmell(w.object, w.kind, &container, w.workloadPosition, constants.K8S_SEC_CAPABILITIES_UNSET),
			)
			continue
		}
		if len(container.SecurityContext.Capabilities.Drop) == 0 || !strings.EqualFold(string(container.SecurityContext.Capabilities.Drop[0]), "ALL") {
			w.kubernetesSmells = append(
				w.kubernetesSmells,
				models.NewKubernetesSmell(w.object, w.kind, &container, w.workloadPosition, constants.K8S_SEC_CAPABILITIES_VALUE),
			)
		}
	}
}

func (w *k8sWorkload) SmellySecurityContextAllowPrivilegeEscalation(spec *corev1.PodSpec) {
	if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
		w.kubernetesSmells = append(
			w.kubernetesSmells,
			models.NewKubernetesSmell(w.object, w.kind, container, w.workloadPosition, constants.K8S_SEC_PRIVESCALATION_UNSET),
		)
		return
	}
	if *container.SecurityContext.AllowPrivilegeEscalation {
		w.kubernetesSmells = append(
			w.kubernetesSmells,
			models.NewKubernetesSmell(w.object, w.kind, container, w.workloadPosition, constants.K8S_SEC_PRIVESCALATION_VALUE),
		)
	}
}

func (w *k8sWorkload) SmellySecurityContextReadOnlyRootFilesystem(spec *corev1.PodSpec) {
	if container.SecurityContext == nil || container.SecurityContext.ReadOnlyRootFilesystem == nil {
		w.kubernetesSmells = append(
			w.kubernetesSmells,
			models.NewKubernetesSmell(w.object, w.kind, container, w.workloadPosition, constants.K8S_SEC_ROROOTFS_UNSET),
		)
		return
	}
	if !*container.SecurityContext.ReadOnlyRootFilesystem {
		w.kubernetesSmells = append(
			w.kubernetesSmells,
			models.NewKubernetesSmell(w.object, w.kind, container, w.workloadPosition, constants.K8S_SEC_ROROOTFS_VALUE),
		)
	}
}

func (w *k8sWorkload) SmellySecurityContextPrivileged(spec *corev1.PodSpec) {
	if container.SecurityContext == nil || container.SecurityContext.Privileged == nil {
		w.kubernetesSmells = append(
			w.kubernetesSmells,
			models.NewKubernetesSmell(w.object, w.kind, container, w.workloadPosition, constants.K8S_SEC_PRIVILEGED_UNSET),
		)
	}
}

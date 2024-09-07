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
	SmellyResourceAndLimit()
	SmellySecurityContextRunAsUser()
	SmellySecurityContextRunAsNonRoot()
	SmellySecurityContextCapabilities()
	SmellySecurityContextAllowPrivilegeEscalation()
	SmellySecurityContextReadOnlyRootFilesystem()
	SmellySecurityContextPrivileged()
}

type k8sWorkload struct {
	object           metav1.Object
	kind             schema.ObjectKind
	workloadPosition int
	kubernetesSmells []*models.KubernetesSmell
}

func NewK8sWorkload(object metav1.Object, kind schema.ObjectKind, workloadPosition int) *k8sWorkload {
	return &k8sWorkload{
		object:           object,
		kind:             kind,
		workloadPosition: workloadPosition,
	}
}

func (w *k8sWorkload) GetKubernetesSmells() []*models.KubernetesSmell {
	return w.kubernetesSmells
}

func (w *k8sWorkload) AddKubernetesSmell(smell *models.KubernetesSmell) {
	w.kubernetesSmells = append(w.kubernetesSmells, smell)
}

func (w *k8sWorkload) SmellySecurityContextRunAsUser(spec *corev1.PodSpec) {
	isSetAtPodLevel := spec.SecurityContext != nil && spec.SecurityContext.RunAsUser != nil
	isValueSmellAtPodLevel := isSetAtPodLevel && *spec.SecurityContext.RunAsUser == 0
	for _, container := range spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			if !isSetAtPodLevel {
				w.kubernetesSmells = append(
					w.kubernetesSmells,
					models.NewKubernetesSmell(w.object, w.kind, &container, w.workloadPosition, constants.K8S_SEC_RUNASUSER_UNSET),
				)
			} else if isValueSmellAtPodLevel {
				w.kubernetesSmells = append(
					w.kubernetesSmells,
					models.NewKubernetesSmell(w.object, w.kind, &container, w.workloadPosition, constants.K8S_SEC_RUNASUSER_VALUE),
				)
			}
			continue
		}
		if *container.SecurityContext.RunAsUser == 0 {
			w.kubernetesSmells = append(
				w.kubernetesSmells,
				models.NewKubernetesSmell(w.object, w.kind, &container, w.workloadPosition, constants.K8S_SEC_RUNASUSER_VALUE),
			)
		}
	}
}

func (w *k8sWorkload) SmellySecurityContextRunAsNonRoot(spec *corev1.PodSpec) {
	isSetAtPodLevel := spec.SecurityContext != nil && spec.SecurityContext.RunAsNonRoot != nil
	isValueSmellAtPodLevel := isSetAtPodLevel && !*spec.SecurityContext.RunAsNonRoot
	for _, container := range spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsNonRoot == nil {
			if !isSetAtPodLevel {
				w.kubernetesSmells = append(
					w.kubernetesSmells,
					models.NewKubernetesSmell(w.object, w.kind, &container, w.workloadPosition, constants.K8S_SEC_RUNASNONROOT_UNSET),
				)
			} else if isValueSmellAtPodLevel {
				w.kubernetesSmells = append(
					w.kubernetesSmells,
					models.NewKubernetesSmell(w.object, w.kind, &container, w.workloadPosition, constants.K8S_SEC_RUNASNONROOT_VALUE),
				)
			}
			continue
		}
		if !*container.SecurityContext.RunAsNonRoot {
			w.kubernetesSmells = append(
				w.kubernetesSmells,
				models.NewKubernetesSmell(w.object, w.kind, &container, w.workloadPosition, constants.K8S_SEC_RUNASNONROOT_VALUE),
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
	for _, container := range spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			w.kubernetesSmells = append(
				w.kubernetesSmells,
				models.NewKubernetesSmell(w.object, w.kind, &container, w.workloadPosition, constants.K8S_SEC_PRIVESCALATION_UNSET),
			)
			continue
		}
		if *container.SecurityContext.AllowPrivilegeEscalation {
			w.kubernetesSmells = append(
				w.kubernetesSmells,
				models.NewKubernetesSmell(w.object, w.kind, &container, w.workloadPosition, constants.K8S_SEC_PRIVESCALATION_VALUE),
			)
		}
	}
}

func (w *k8sWorkload) SmellySecurityContextReadOnlyRootFilesystem(spec *corev1.PodSpec) {
	for _, container := range spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.ReadOnlyRootFilesystem == nil {
			w.kubernetesSmells = append(
				w.kubernetesSmells,
				models.NewKubernetesSmell(w.object, w.kind, &container, w.workloadPosition, constants.K8S_SEC_ROROOTFS_UNSET),
			)
			continue
		}
		if !*container.SecurityContext.ReadOnlyRootFilesystem {
			w.kubernetesSmells = append(
				w.kubernetesSmells,
				models.NewKubernetesSmell(w.object, w.kind, &container, w.workloadPosition, constants.K8S_SEC_ROROOTFS_VALUE),
			)
		}
	}
}

func (w *k8sWorkload) SmellySecurityContextPrivileged(spec *corev1.PodSpec) {
	for _, container := range spec.Containers {
		if container.SecurityContext != nil && container.SecurityContext.Privileged != nil && *container.SecurityContext.Privileged {
			w.kubernetesSmells = append(
				w.kubernetesSmells,
				models.NewKubernetesSmell(w.object, w.kind, &container, w.workloadPosition, constants.K8S_SEC_PRIVILEGED_VALUE),
			)
		}
	}
}

func (w *k8sWorkload) SmellyResourceAndLimit(spec *corev1.PodSpec) {
	for _, container := range spec.Containers {
		if container.Resources.Requests == nil {
			w.kubernetesSmells = append(
				w.kubernetesSmells,
				models.NewKubernetesSmell(w.object, w.kind, &container, w.workloadPosition, constants.K8S_SEC_RESREQUESTS_UNSET),
			)
		}
		if container.Resources.Limits == nil {
			w.kubernetesSmells = append(
				w.kubernetesSmells,
				models.NewKubernetesSmell(w.object, w.kind, &container, w.workloadPosition, constants.K8S_SEC_RESLIMITS_UNSET),
			)
		}
	}
}

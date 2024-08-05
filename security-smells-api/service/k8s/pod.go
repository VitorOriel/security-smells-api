package k8s

import (
	corev1 "k8s.io/api/core/v1"
)

type Pod interface {
	K8sWorkload
}

type pod struct {
	*k8sWorkload
	Pod *corev1.Pod
}

func NewPod(object *corev1.Pod, workloadPosition int) Pod {
	return &pod{
		k8sWorkload: NewK8sWorkload(object, object.GetObjectKind(), workloadPosition),
		Pod:         object,
	}
}

func (p *pod) SmellySecurityContextReadOnlyRootFilesystem() {
	p.k8sWorkload.SmellySecurityContextReadOnlyRootFilesystem(&p.Pod.Spec)
}

func (p *pod) SmellySecurityContextAllowPrivilegeEscalation() {
	p.k8sWorkload.SmellySecurityContextAllowPrivilegeEscalation(&p.Pod.Spec)
}

func (p *pod) SmellySecurityContextCapabilities() {
	p.k8sWorkload.SmellySecurityContextCapabilities(&p.Pod.Spec)
}

func (p *pod) SmellySecurityContextPrivileged() {
	p.k8sWorkload.SmellySecurityContextPrivileged(&p.Pod.Spec)
}

func (p *pod) SmellySecurityContextRunAsUser() {
	p.k8sWorkload.SmellySecurityContextRunAsUser(&p.Pod.Spec)
}

func (p *pod) SmellyResourceAndLimit() {
	p.k8sWorkload.SmellyResourceAndLimit(&p.Pod.Spec)
}

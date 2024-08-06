package k8s

import (
	appsv1 "k8s.io/api/apps/v1"
)

type DaemonSet interface {
	K8sWorkload
}

type daemonSet struct {
	*k8sWorkload
	DaemonSet *appsv1.DaemonSet
}

func NewDaemonSet(object *appsv1.DaemonSet, workloadPosition int) DaemonSet {
	return &daemonSet{
		k8sWorkload: NewK8sWorkload(object, object.GetObjectKind(), workloadPosition),
		DaemonSet:   object,
	}
}

func (d *daemonSet) SmellySecurityContextReadOnlyRootFilesystem() {
	d.k8sWorkload.SmellySecurityContextReadOnlyRootFilesystem(&d.DaemonSet.Spec.Template.Spec)
}

func (d *daemonSet) SmellySecurityContextAllowPrivilegeEscalation() {
	d.k8sWorkload.SmellySecurityContextAllowPrivilegeEscalation(&d.DaemonSet.Spec.Template.Spec)
}

func (d *daemonSet) SmellySecurityContextCapabilities() {
	d.k8sWorkload.SmellySecurityContextCapabilities(&d.DaemonSet.Spec.Template.Spec)
}

func (d *daemonSet) SmellySecurityContextPrivileged() {
	d.k8sWorkload.SmellySecurityContextPrivileged(&d.DaemonSet.Spec.Template.Spec)
}

func (d *daemonSet) SmellySecurityContextRunAsUser() {
	d.k8sWorkload.SmellySecurityContextRunAsUser(&d.DaemonSet.Spec.Template.Spec)
}

func (d *daemonSet) SmellySecurityContextRunAsNonRoot() {
	d.k8sWorkload.SmellySecurityContextRunAsNonRoot(&d.DaemonSet.Spec.Template.Spec)
}

func (d *daemonSet) SmellyResourceAndLimit() {
	d.k8sWorkload.SmellyResourceAndLimit(&d.DaemonSet.Spec.Template.Spec)
}

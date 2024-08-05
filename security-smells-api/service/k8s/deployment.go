package k8s

import (
	appsv1 "k8s.io/api/apps/v1"
)

type Deployment interface {
	K8sWorkload
}

type deployment struct {
	*k8sWorkload
	Deployment *appsv1.Deployment
}

func NewDeployment(object *appsv1.Deployment, workloadPosition int) Deployment {
	return &deployment{
		k8sWorkload: NewK8sWorkload(object, object.GetObjectKind(), workloadPosition),
		Deployment:  object,
	}
}

func (d *deployment) SmellySecurityContextReadOnlyRootFilesystem() {
	d.k8sWorkload.SmellySecurityContextReadOnlyRootFilesystem(&d.Deployment.Spec.Template.Spec)
}

func (d *deployment) SmellySecurityContextAllowPrivilegeEscalation() {
	d.k8sWorkload.SmellySecurityContextAllowPrivilegeEscalation(&d.Deployment.Spec.Template.Spec)
}

func (d *deployment) SmellySecurityContextCapabilities() {
	d.k8sWorkload.SmellySecurityContextCapabilities(&d.Deployment.Spec.Template.Spec)
}

func (d *deployment) SmellySecurityContextPrivileged() {
	d.k8sWorkload.SmellySecurityContextPrivileged(&d.Deployment.Spec.Template.Spec)
}

func (d *deployment) SmellySecurityContextRunAsUser() {
	d.k8sWorkload.SmellySecurityContextRunAsUser(&d.Deployment.Spec.Template.Spec)
}

func (d *deployment) SmellyResourceAndLimit() {
	d.k8sWorkload.SmellyResourceAndLimit(&d.Deployment.Spec.Template.Spec)
}

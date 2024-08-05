package k8s

import (
	v1 "k8s.io/api/apps/v1"
)

type StatefulSet interface {
	K8sWorkload
}

type statefulSet struct {
	*k8sWorkload
	StatefulSet *v1.StatefulSet
}

func NewStatefulSet(object *v1.StatefulSet, workloadPosition int) StatefulSet {
	return &statefulSet{
		k8sWorkload: NewK8sWorkload(object, object.GetObjectKind(), workloadPosition),
		StatefulSet: object,
	}
}

func (s *statefulSet) SmellySecurityContextReadOnlyRootFilesystem() {
	s.k8sWorkload.SmellySecurityContextReadOnlyRootFilesystem(&s.StatefulSet.Spec.Template.Spec)
}

func (s *statefulSet) SmellySecurityContextAllowPrivilegeEscalation() {
	s.k8sWorkload.SmellySecurityContextAllowPrivilegeEscalation(&s.StatefulSet.Spec.Template.Spec)
}

func (s *statefulSet) SmellySecurityContextCapabilities() {
	s.k8sWorkload.SmellySecurityContextCapabilities(&s.StatefulSet.Spec.Template.Spec)
}

func (s *statefulSet) SmellySecurityContextPrivileged() {
	s.k8sWorkload.SmellySecurityContextPrivileged(&s.StatefulSet.Spec.Template.Spec)
}

func (s *statefulSet) SmellySecurityContextRunAsUser() {
	s.k8sWorkload.SmellySecurityContextRunAsUser(&s.StatefulSet.Spec.Template.Spec)
}

func (s *statefulSet) SmellyResourceAndLimit() {
	s.k8sWorkload.SmellyResourceAndLimit(&s.StatefulSet.Spec.Template.Spec)
}

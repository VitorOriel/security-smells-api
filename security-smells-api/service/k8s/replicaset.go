package k8s

import (
	appsv1 "k8s.io/api/apps/v1"
)

// type ReplicaSet struct {
// 	ReplicaSet       *appsv1.ReplicaSet

type ReplicaSet interface {
	K8sWorkload
}

type replicaSet struct {
	*k8sWorkload
	ReplicaSet *appsv1.ReplicaSet
}

func NewReplicaSet(object *appsv1.ReplicaSet, workloadPosition int) ReplicaSet {
	return &replicaSet{
		k8sWorkload: NewK8sWorkload(object, object.GetObjectKind(), workloadPosition),
		ReplicaSet:  object,
	}
}

func (r *replicaSet) SmellySecurityContextReadOnlyRootFilesystem() {
	r.k8sWorkload.SmellySecurityContextReadOnlyRootFilesystem(&r.ReplicaSet.Spec.Template.Spec)
}

func (r *replicaSet) SmellySecurityContextAllowPrivilegeEscalation() {
	r.k8sWorkload.SmellySecurityContextAllowPrivilegeEscalation(&r.ReplicaSet.Spec.Template.Spec)
}

func (r *replicaSet) SmellySecurityContextCapabilities() {
	r.k8sWorkload.SmellySecurityContextCapabilities(&r.ReplicaSet.Spec.Template.Spec)
}

func (r *replicaSet) SmellySecurityContextPrivileged() {
	r.k8sWorkload.SmellySecurityContextPrivileged(&r.ReplicaSet.Spec.Template.Spec)
}

func (r *replicaSet) SmellySecurityContextRunAsUser() {
	r.k8sWorkload.SmellySecurityContextRunAsUser(&r.ReplicaSet.Spec.Template.Spec)
}

func (r *replicaSet) SmellyResourceAndLimit() {
	r.k8sWorkload.SmellyResourceAndLimit(&r.ReplicaSet.Spec.Template.Spec)
}

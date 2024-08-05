package k8s

import (
	batchv1 "k8s.io/api/batch/v1"
)

// type Job struct {
// 	Job              *batchv1.Job

type Job interface {
	K8sWorkload
}

type job struct {
	*k8sWorkload
	Job *batchv1.Job
}

func NewJob(object *batchv1.Job, workloadPosition int) Job {
	return &job{
		k8sWorkload: NewK8sWorkload(object, object.GetObjectKind(), workloadPosition),
		Job:         object,
	}
}

func (j *job) SmellySecurityContextReadOnlyRootFilesystem() {
	j.k8sWorkload.SmellySecurityContextReadOnlyRootFilesystem(&j.Job.Spec.Template.Spec)
}

func (j *job) SmellySecurityContextAllowPrivilegeEscalation() {
	j.k8sWorkload.SmellySecurityContextAllowPrivilegeEscalation(&j.Job.Spec.Template.Spec)
}

func (j *job) SmellySecurityContextCapabilities() {
	j.k8sWorkload.SmellySecurityContextCapabilities(&j.Job.Spec.Template.Spec)
}

func (j *job) SmellySecurityContextPrivileged() {
	j.k8sWorkload.SmellySecurityContextPrivileged(&j.Job.Spec.Template.Spec)
}

func (j *job) SmellySecurityContextRunAsUser() {
	j.k8sWorkload.SmellySecurityContextRunAsUser(&j.Job.Spec.Template.Spec)
}

func (j *job) SmellyResourceAndLimit() {
	j.k8sWorkload.SmellyResourceAndLimit(&j.Job.Spec.Template.Spec)
}

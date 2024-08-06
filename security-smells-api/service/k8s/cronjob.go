package k8s

import (
	batchv1 "k8s.io/api/batch/v1"
)

type CronJob interface {
	K8sWorkload
}

type cronJob struct {
	*k8sWorkload
	CronJob *batchv1.CronJob
}

func NewCronJob(object *batchv1.CronJob, workloadPosition int) CronJob {
	return &cronJob{
		k8sWorkload: NewK8sWorkload(object, object.GetObjectKind(), workloadPosition),
		CronJob:     object,
	}
}

func (c *cronJob) SmellySecurityContextReadOnlyRootFilesystem() {
	c.k8sWorkload.SmellySecurityContextReadOnlyRootFilesystem(&c.CronJob.Spec.JobTemplate.Spec.Template.Spec)
}

func (c *cronJob) SmellySecurityContextAllowPrivilegeEscalation() {
	c.k8sWorkload.SmellySecurityContextAllowPrivilegeEscalation(&c.CronJob.Spec.JobTemplate.Spec.Template.Spec)
}

func (c *cronJob) SmellySecurityContextCapabilities() {
	c.k8sWorkload.SmellySecurityContextCapabilities(&c.CronJob.Spec.JobTemplate.Spec.Template.Spec)
}

func (c *cronJob) SmellySecurityContextPrivileged() {
	c.k8sWorkload.SmellySecurityContextPrivileged(&c.CronJob.Spec.JobTemplate.Spec.Template.Spec)
}

func (c *cronJob) SmellySecurityContextRunAsUser() {
	c.k8sWorkload.SmellySecurityContextRunAsUser(&c.CronJob.Spec.JobTemplate.Spec.Template.Spec)
}

func (c *cronJob) SmellySecurityContextRunAsNonRoot() {
	c.k8sWorkload.SmellySecurityContextRunAsNonRoot(&c.CronJob.Spec.JobTemplate.Spec.Template.Spec)
}

func (c *cronJob) SmellyResourceAndLimit() {
	c.k8sWorkload.SmellyResourceAndLimit(&c.CronJob.Spec.JobTemplate.Spec.Template.Spec)
}

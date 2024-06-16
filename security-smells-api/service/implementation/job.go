package implementation

import (
	"security-smells-api/constants"
	"security-smells-api/models"
	"security-smells-api/service/interfaces"

	batchv1 "k8s.io/api/batch/v1"
)

type Job struct {
	interfaces.SmellyJob
	Job              *batchv1.Job
	WorkloadPosition int
	SmellKubernetes  []*models.SmellKubernetes
}

func (job *Job) SmellySecurityContextAllowPrivilegeEscalation() {
	for _, container := range job.Job.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			job.SmellKubernetes = append(job.SmellKubernetes, models.NewSmellKubernetes(job.Job, job.Job.GetObjectKind(), &container, job.WorkloadPosition, constants.K8S_SEC_PRIVESCALATION))
		}
	}
}

func (job *Job) SmellySecurityContextCapabilities() {
	for _, container := range job.Job.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			job.SmellKubernetes = append(job.SmellKubernetes, models.NewSmellKubernetes(job.Job, job.Job.GetObjectKind(), &container, job.WorkloadPosition, constants.K8S_SEC_CAPABILITIES))
		}
	}
}

func (job *Job) SmellySecurityContextRunAsUser() {
	for _, container := range job.Job.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			job.SmellKubernetes = append(job.SmellKubernetes, models.NewSmellKubernetes(job.Job, job.Job.GetObjectKind(), &container, job.WorkloadPosition, constants.K8S_SEC_RUNASUSER))
		}
	}
}

func (job *Job) SmellyResourceAndLimit() {
	for _, container := range job.Job.Spec.Template.Spec.Containers {
		if container.Resources.Requests == nil {
			job.SmellKubernetes = append(job.SmellKubernetes, models.NewSmellKubernetes(job.Job, job.Job.GetObjectKind(), &container, job.WorkloadPosition, constants.K8S_SEC_RESREQUESTS))
		}
		if container.Resources.Limits == nil {
			job.SmellKubernetes = append(job.SmellKubernetes, models.NewSmellKubernetes(job.Job, job.Job.GetObjectKind(), &container, job.WorkloadPosition, constants.K8S_SEC_RESLIMITS))
		}
	}
}

func (job *Job) SmellySecurityContextReadOnlyRootFilesystem() {
	for _, container := range job.Job.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.ReadOnlyRootFilesystem == nil {
			job.SmellKubernetes = append(job.SmellKubernetes, models.NewSmellKubernetes(job.Job, job.Job.GetObjectKind(), &container, job.WorkloadPosition, constants.K8S_SEC_ROROOTFS))
		}
	}
}

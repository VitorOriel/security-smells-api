package k8s

import (
	"security-smells-api/constants"
	"security-smells-api/models"

	batchv1 "k8s.io/api/batch/v1"
)

type Job struct {
	Job              *batchv1.Job
	WorkloadPosition int
	KubernetesSmell  []*models.KubernetesSmell
}

func (job *Job) SmellySecurityContextAllowPrivilegeEscalation() {
	for _, container := range job.Job.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			job.KubernetesSmell = append(job.KubernetesSmell, models.NewKubernetesSmell(job.Job, job.Job.GetObjectKind(), &container, job.WorkloadPosition, constants.K8S_SEC_PRIVESCALATION_UNSET))
		}
	}
}

func (job *Job) SmellySecurityContextCapabilities() {
	for _, container := range job.Job.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			job.KubernetesSmell = append(job.KubernetesSmell, models.NewKubernetesSmell(job.Job, job.Job.GetObjectKind(), &container, job.WorkloadPosition, constants.K8S_SEC_CAPABILITIES_UNSET))
		}
	}
}

func (job *Job) SmellySecurityContextRunAsUser() {
	for _, container := range job.Job.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			job.KubernetesSmell = append(job.KubernetesSmell, models.NewKubernetesSmell(job.Job, job.Job.GetObjectKind(), &container, job.WorkloadPosition, constants.K8S_SEC_RUNASUSER_UNSET))
		}
	}
}

func (job *Job) SmellyResourceAndLimit() {
	for _, container := range job.Job.Spec.Template.Spec.Containers {
		if container.Resources.Requests == nil {
			job.KubernetesSmell = append(job.KubernetesSmell, models.NewKubernetesSmell(job.Job, job.Job.GetObjectKind(), &container, job.WorkloadPosition, constants.K8S_SEC_RESREQUESTS_UNSET))
		}
		if container.Resources.Limits == nil {
			job.KubernetesSmell = append(job.KubernetesSmell, models.NewKubernetesSmell(job.Job, job.Job.GetObjectKind(), &container, job.WorkloadPosition, constants.K8S_SEC_RESLIMITS_UNSET))
		}
	}
}

func (job *Job) SmellySecurityContextReadOnlyRootFilesystem() {
	for _, container := range job.Job.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.ReadOnlyRootFilesystem == nil {
			job.KubernetesSmell = append(job.KubernetesSmell, models.NewKubernetesSmell(job.Job, job.Job.GetObjectKind(), &container, job.WorkloadPosition, constants.K8S_SEC_ROROOTFS_UNSET))
		}
	}
}

func (job *Job) SmellySecurityContextPrivileged() {
	for _, container := range job.Job.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Privileged == nil {
			job.KubernetesSmell = append(job.KubernetesSmell, models.NewKubernetesSmell(job.Job, job.Job.GetObjectKind(), &container, job.WorkloadPosition, constants.K8S_SEC_PRIVILEGED_UNSET))
		}
	}
}

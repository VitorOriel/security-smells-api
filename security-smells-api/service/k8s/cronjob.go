package k8s

import (
	"security-smells-api/constants"
	"security-smells-api/models"

	batchv1 "k8s.io/api/batch/v1"
)

type CronJob struct {
	CronJob          *batchv1.CronJob
	WorkloadPosition int
	KubernetesSmell  []*models.KubernetesSmell
}

func (cronJob *CronJob) SmellySecurityContextAllowPrivilegeEscalation() {
	for _, container := range cronJob.CronJob.Spec.JobTemplate.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			cronJob.KubernetesSmell = append(cronJob.KubernetesSmell, models.NewKubernetesSmell(cronJob.CronJob, cronJob.CronJob.GetObjectKind(), &container, cronJob.WorkloadPosition, constants.K8S_SEC_PRIVESCALATION_UNSET))
		}
	}
}

func (cronJob *CronJob) SmellySecurityContextCapabilities() {
	for _, container := range cronJob.CronJob.Spec.JobTemplate.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			cronJob.KubernetesSmell = append(cronJob.KubernetesSmell, models.NewKubernetesSmell(cronJob.CronJob, cronJob.CronJob.GetObjectKind(), &container, cronJob.WorkloadPosition, constants.K8S_SEC_CAPABILITIES_UNSET))
		}
	}
}

func (cronJob *CronJob) SmellySecurityContextRunAsUser() {
	for _, container := range cronJob.CronJob.Spec.JobTemplate.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			cronJob.KubernetesSmell = append(cronJob.KubernetesSmell, models.NewKubernetesSmell(cronJob.CronJob, cronJob.CronJob.GetObjectKind(), &container, cronJob.WorkloadPosition, constants.K8S_SEC_RUNASUSER_UNSET))
		}
	}
}

func (cronJob *CronJob) SmellyResourceAndLimit() {
	for _, container := range cronJob.CronJob.Spec.JobTemplate.Spec.Template.Spec.Containers {
		if container.Resources.Requests == nil {
			cronJob.KubernetesSmell = append(cronJob.KubernetesSmell, models.NewKubernetesSmell(cronJob.CronJob, cronJob.CronJob.GetObjectKind(), &container, cronJob.WorkloadPosition, constants.K8S_SEC_RESREQUESTS_UNSET))
		}
		if container.Resources.Limits == nil {
			cronJob.KubernetesSmell = append(cronJob.KubernetesSmell, models.NewKubernetesSmell(cronJob.CronJob, cronJob.CronJob.GetObjectKind(), &container, cronJob.WorkloadPosition, constants.K8S_SEC_RESLIMITS_UNSET))
		}
	}
}

func (cronJob *CronJob) SmellySecurityContextReadOnlyRootFilesystem() {
	for _, container := range cronJob.CronJob.Spec.JobTemplate.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.ReadOnlyRootFilesystem == nil {
			cronJob.KubernetesSmell = append(cronJob.KubernetesSmell, models.NewKubernetesSmell(cronJob.CronJob, cronJob.CronJob.GetObjectKind(), &container, cronJob.WorkloadPosition, constants.K8S_SEC_ROROOTFS_UNSET))
		}
	}
}

func (cronJob *CronJob) SmellySecurityContextPrivileged() {
	for _, container := range cronJob.CronJob.Spec.JobTemplate.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Privileged == nil {
			cronJob.KubernetesSmell = append(cronJob.KubernetesSmell, models.NewKubernetesSmell(cronJob.CronJob, cronJob.CronJob.GetObjectKind(), &container, cronJob.WorkloadPosition, constants.K8S_SEC_PRIVILEGED_UNSET))
		}
	}
}

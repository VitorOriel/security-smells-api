package implementation

import (
	"security-smells-api/models"
	"security-smells-api/service/interfaces"

	batchv1 "k8s.io/api/batch/v1"
)

type CronJob struct {
	interfaces.SmellyDeployment
	CronJob      *batchv1.CronJob
	SmellCronJob []models.SmellCronJob
}

func (cronJob *CronJob) SmellySecurityContextAllowPrivilegeEscalation() {
	nameSpace := cronJob.CronJob.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	cronJobName := cronJob.CronJob.GetName()
	kind := cronJob.CronJob.GroupVersionKind().Kind
	for _, container := range cronJob.CronJob.Spec.JobTemplate.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			smellCronJob := models.SmellCronJob{
				NameSpace:      nameSpace,
				CronJobName:    cronJobName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "AllowPrivilegeEscalation not set into " + container.Name + " your container is running with AllowPrivilegeEscalation",
				Suggestion:     "Please add AllowPrivilegeEscalation into " + container.Name + " to avoid running with AllowPrivilegeEscalation",
			}
			cronJob.SmellCronJob = append(cronJob.SmellCronJob, smellCronJob)
		}
	}
}

func (cronJob *CronJob) SmellySecurityContextCapabilities() {
	nameSpace := cronJob.CronJob.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	cronJobName := cronJob.CronJob.GetName()
	kind := cronJob.CronJob.GroupVersionKind().Kind
	for _, container := range cronJob.CronJob.Spec.JobTemplate.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			smellCronJob := models.SmellCronJob{
				NameSpace:      nameSpace,
				CronJobName:    cronJobName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "Capabilities not set into " + container.Name + " your container is running without Capabilities",
				Suggestion:     "Please add Capabilities into " + container.Name + " to avoid running without Capabilities",
			}
			cronJob.SmellCronJob = append(cronJob.SmellCronJob, smellCronJob)
		}
	}
}

func (cronJob *CronJob) SmellySecurityContextRunAsUser() {
	nameSpace := cronJob.CronJob.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	cronJobName := cronJob.CronJob.GetName()
	kind := cronJob.CronJob.GroupVersionKind().Kind
	for _, container := range cronJob.CronJob.Spec.JobTemplate.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			smellCronJob := models.SmellCronJob{
				NameSpace:      nameSpace,
				CronJobName:    cronJobName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "RunAsUser not set into " + container.Name + " your container is running with root user",
				Suggestion:     "Please add RunAsUser into " + container.Name + " to avoid running with root user",
			}
			cronJob.SmellCronJob = append(cronJob.SmellCronJob, smellCronJob)
		}
	}
}

func (CronJob *CronJob) SmellyResourceAndLimit() {
	nameSpace := CronJob.CronJob.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	CronJobName := CronJob.CronJob.GetName()
	kind := CronJob.CronJob.GroupVersionKind().Kind
	for _, container := range CronJob.CronJob.Spec.JobTemplate.Spec.Template.Spec.Containers {
		if container.Resources.Requests == nil {
			smellCronJob := models.SmellCronJob{
				NameSpace:      nameSpace,
				CronJobName:    CronJobName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "Resources not set into " + container.Name + " your container is running without Resources",
				Suggestion:     "Please add Resources into " + container.Name + " to avoid running without Resources",
			}
			CronJob.SmellCronJob = append(CronJob.SmellCronJob, smellCronJob)
		}
		if container.Resources.Limits == nil {
			smellCronJob := models.SmellCronJob{
				NameSpace:      nameSpace,
				CronJobName:    CronJobName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "Limits not set into " + container.Name + " your container is running without Limits",
				Suggestion:     "Please add Limits into " + container.Name + " to avoid running without Limits",
			}
			CronJob.SmellCronJob = append(CronJob.SmellCronJob, smellCronJob)
		}
	}
}

func (CronJob *CronJob) SmellySecurityContextReadOnlyRootFilesystem() {
	nameSpace := CronJob.CronJob.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	CronJobName := CronJob.CronJob.GetName()
	kind := CronJob.CronJob.GroupVersionKind().Kind
	for _, container := range CronJob.CronJob.Spec.JobTemplate.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.ReadOnlyRootFilesystem == nil {
			smellCronJob := models.SmellCronJob{
				NameSpace:      nameSpace,
				CronJobName:    CronJobName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "ReadOnlyRootFilesystem not set into " + container.Name + " your container is running with ReadWriteRootFilesystem",
				Suggestion:     "Please add ReadOnlyRootFilesystem into " + container.Name + " to avoid running with ReadWriteRootFilesystem",
			}
			CronJob.SmellCronJob = append(CronJob.SmellCronJob, smellCronJob)
		}
	}
}

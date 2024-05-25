package implementation

import (
	"security-smells-api/models"
	"security-smells-api/service/interfaces"

	batchv1 "k8s.io/api/batch/v1"
)

type CronJob struct {
	interfaces.SmellyDeployment
	CronJob          *batchv1.CronJob
	WorkloadPosition int
	SmellKubernetes  []models.SmellKubernetes
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
			smellCronJob := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: cronJobName,
				WorkloadPosition:  cronJob.WorkloadPosition,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "AllowPrivilegeEscalation not set into " + container.Name + " your container is running with AllowPrivilegeEscalation",
				Suggestion:        "Please add AllowPrivilegeEscalation into " + container.Name + " to avoid running with AllowPrivilegeEscalation",
			}
			cronJob.SmellKubernetes = append(cronJob.SmellKubernetes, smellCronJob)
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
			smellCronJob := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: cronJobName,
				WorkloadPosition:  cronJob.WorkloadPosition,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "Capabilities not set into " + container.Name + " your container is running without Capabilities",
				Suggestion:        "Please add Capabilities into " + container.Name + " to avoid running without Capabilities",
			}
			cronJob.SmellKubernetes = append(cronJob.SmellKubernetes, smellCronJob)
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
			smellCronJob := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: cronJobName,
				WorkloadPosition:  cronJob.WorkloadPosition,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "RunAsUser not set into " + container.Name + " your container is running with root user",
				Suggestion:        "Please add RunAsUser into " + container.Name + " to avoid running with root user",
			}
			cronJob.SmellKubernetes = append(cronJob.SmellKubernetes, smellCronJob)
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
			smellCronJob := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: CronJobName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "Resources not set into " + container.Name + " your container is running without Resources",
				Suggestion:        "Please add Resources into " + container.Name + " to avoid running without Resources",
			}
			CronJob.SmellKubernetes = append(CronJob.SmellKubernetes, smellCronJob)
		}
		if container.Resources.Limits == nil {
			smellCronJob := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: CronJobName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "Limits not set into " + container.Name + " your container is running without Limits",
				Suggestion:        "Please add Limits into " + container.Name + " to avoid running without Limits",
			}
			CronJob.SmellKubernetes = append(CronJob.SmellKubernetes, smellCronJob)
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
			smellCronJob := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: CronJobName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "ReadOnlyRootFilesystem not set into " + container.Name + " your container is running with ReadWriteRootFilesystem",
				Suggestion:        "Please add ReadOnlyRootFilesystem into " + container.Name + " to avoid running with ReadWriteRootFilesystem",
			}
			CronJob.SmellKubernetes = append(CronJob.SmellKubernetes, smellCronJob)
		}
	}
}

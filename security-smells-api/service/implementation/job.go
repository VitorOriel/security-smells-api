package implementation

import (
	"security-smells-api/models"
	"security-smells-api/service/interfaces"

	batchv1 "k8s.io/api/batch/v1"
)

type Job struct {
	interfaces.SmellyJob
	Job      *batchv1.Job
	SmellJob []models.SmellJob
}

func (job *Job) SmellySecurityContextAllowPrivilegeEscalation() {
	nameSpace := job.Job.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	jobName := job.Job.GetName()
	kind := job.Job.GroupVersionKind().Kind
	for _, container := range job.Job.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.AllowPrivilegeEscalation == nil {
			smellJob := models.SmellJob{
				NameSpace:      nameSpace,
				JobName:        jobName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "AllowPrivilegeEscalation not set into " + container.Name + " your container is running with AllowPrivilegeEscalation",
				Suggestion:     "Please add AllowPrivilegeEscalation into " + container.Name + " to avoid running with AllowPrivilegeEscalation",
			}
			job.SmellJob = append(job.SmellJob, smellJob)
		}
	}
}

func (job *Job) SmellySecurityContextCapabilities() {
	nameSpace := job.Job.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	jobName := job.Job.GetName()
	kind := job.Job.GroupVersionKind().Kind
	for _, container := range job.Job.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.Capabilities == nil {
			smellJob := models.SmellJob{
				NameSpace:      nameSpace,
				JobName:        jobName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "Capabilities not set into " + container.Name + " your container is running without Capabilities",
				Suggestion:     "Please add Capabilities into " + container.Name + " to avoid running without Capabilities",
			}
			job.SmellJob = append(job.SmellJob, smellJob)
		}
	}
}

func (job *Job) SmellySecurityContextRunAsUser() {
	nameSpace := job.Job.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	jobName := job.Job.GetName()
	kind := job.Job.GroupVersionKind().Kind
	for _, container := range job.Job.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.RunAsUser == nil {
			smellJob := models.SmellJob{
				NameSpace:      nameSpace,
				JobName:        jobName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "RunAsUser not set into " + container.Name + " your container is running with root user",
				Suggestion:     "Please add RunAsUser into " + container.Name + " to avoid running with root user",
			}
			job.SmellJob = append(job.SmellJob, smellJob)
		}
	}
}

func (Job *Job) SmellyResourceAndLimit() {
	nameSpace := Job.Job.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	JobName := Job.Job.GetName()
	kind := Job.Job.GroupVersionKind().Kind
	for _, container := range Job.Job.Spec.Template.Spec.Containers {
		if container.Resources.Requests == nil {
			smellJob := models.SmellJob{
				NameSpace:      nameSpace,
				JobName:        JobName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "Resources not set into " + container.Name + " your container is running without Resources",
				Suggestion:     "Please add Resources into " + container.Name + " to avoid running without Resources",
			}
			Job.SmellJob = append(Job.SmellJob, smellJob)
		}
		if container.Resources.Limits == nil {
			smellJob := models.SmellJob{
				NameSpace:      nameSpace,
				JobName:        JobName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "Limits not set into " + container.Name + " your container is running without Limits",
				Suggestion:     "Please add Limits into " + container.Name + " to avoid running without Limits",
			}
			Job.SmellJob = append(Job.SmellJob, smellJob)
		}
	}
}

func (Job *Job) SmellySecurityContextReadOnlyRootFilesystem() {
	nameSpace := Job.Job.GetNamespace()
	if nameSpace == "" {
		nameSpace = "default"
	}
	JobName := Job.Job.GetName()
	kind := Job.Job.GroupVersionKind().Kind
	for _, container := range Job.Job.Spec.Template.Spec.Containers {
		if container.SecurityContext == nil || container.SecurityContext.ReadOnlyRootFilesystem == nil {
			smellJob := models.SmellJob{
				NameSpace:      nameSpace,
				JobName:        JobName,
				ContainerName:  container.Name,
				ContainerImage: container.Image,
				Kind:           kind,
				Message:        "ReadOnlyRootFilesystem not set into " + container.Name + " your container is running with ReadWriteRootFilesystem",
				Suggestion:     "Please add ReadOnlyRootFilesystem into " + container.Name + " to avoid running with ReadWriteRootFilesystem",
			}
			Job.SmellJob = append(Job.SmellJob, smellJob)
		}
	}
}

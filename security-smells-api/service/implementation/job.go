package implementation

import (
	"security-smells-api/models"
	"security-smells-api/service/interfaces"

	batchv1 "k8s.io/api/batch/v1"
)

type Job struct {
	interfaces.SmellyJob
	Job             *batchv1.Job
	SmellKubernetes []models.SmellKubernetes
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
			smellJob := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: jobName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "AllowPrivilegeEscalation not set into " + container.Name + " your container is running with AllowPrivilegeEscalation",
				Suggestion:        "Please add AllowPrivilegeEscalation into " + container.Name + " to avoid running with AllowPrivilegeEscalation",
			}
			job.SmellKubernetes = append(job.SmellKubernetes, smellJob)
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
			smellJob := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: jobName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "Capabilities not set into " + container.Name + " your container is running without Capabilities",
				Suggestion:        "Please add Capabilities into " + container.Name + " to avoid running without Capabilities",
			}
			job.SmellKubernetes = append(job.SmellKubernetes, smellJob)
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
			smellJob := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: jobName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "RunAsUser not set into " + container.Name + " your container is running with root user",
				Suggestion:        "Please add RunAsUser into " + container.Name + " to avoid running with root user",
			}
			job.SmellKubernetes = append(job.SmellKubernetes, smellJob)
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
			smellJob := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: JobName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "Resources not set into " + container.Name + " your container is running without Resources",
				Suggestion:        "Please add Resources into " + container.Name + " to avoid running without Resources",
			}
			Job.SmellKubernetes = append(Job.SmellKubernetes, smellJob)
		}
		if container.Resources.Limits == nil {
			smellJob := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: JobName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "Limits not set into " + container.Name + " your container is running without Limits",
				Suggestion:        "Please add Limits into " + container.Name + " to avoid running without Limits",
			}
			Job.SmellKubernetes = append(Job.SmellKubernetes, smellJob)
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
			smellJob := models.SmellKubernetes{
				Namespace:         nameSpace,
				WorkloadKind:      kind,
				WorkloadLabelName: JobName,
				ContainerName:     container.Name,
				ContainerImage:    container.Image,
				Message:           "ReadOnlyRootFilesystem not set into " + container.Name + " your container is running with ReadWriteRootFilesystem",
				Suggestion:        "Please add ReadOnlyRootFilesystem into " + container.Name + " to avoid running with ReadWriteRootFilesystem",
			}
			Job.SmellKubernetes = append(Job.SmellKubernetes, smellJob)
		}
	}
}

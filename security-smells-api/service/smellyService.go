package service

import (
	"errors"
	"security-smells-api/models"
	"security-smells-api/repository"
	"security-smells-api/service/implementation"
	"strings"

	"github.com/gofiber/fiber/v2/log"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

type SmellyService struct {
	SmellyRepository repository.SmellyRepository
}

func (smellyService SmellyService) FindDeploymentSmell(deployments []appsv1.Deployment) (smells []models.SmellDeployment) {
	smells = []models.SmellDeployment{}
	for _, deployment := range deployments {
		d := &implementation.Deployment{
			Deployment: &deployment,
		}
		d.SmellyResourceAndLimit()
		d.SmellySecurityContextRunAsUser()
		d.SmellySecurityContextCapabilities()
		d.SmellySecurityContextAllowPrivilegeEscalation()
		d.SmellySecurityContextReadOnlyRootFilesystem()
		smells = append(smells, d.SmellDeployment...)
	}

	return smells
}

func (smellyService SmellyService) FindPodSmell(pods []corev1.Pod) (smells []models.SmellPod) {
	smells = []models.SmellPod{}
	for _, pod := range pods {
		p := &implementation.Pod{
			Pod: &pod,
		}
		p.SmellyResourceAndLimit()
		p.SmellySecurityContextRunAsUser()
		p.SmellySecurityContextCapabilities()
		p.SmellySecurityContextAllowPrivilegeEscalation()
		p.SmellySecurityContextReadOnlyRootFilesystem()
		smells = append(smells, p.SmellPod...)
	}
	return smells
}

func (smellyService SmellyService) FindJobSmell(jobs []batchv1.Job) (smells []models.SmellJob) {
	smells = []models.SmellJob{}
	for _, job := range jobs {
		j := &implementation.Job{
			Job: &job,
		}
		j.SmellyResourceAndLimit()
		j.SmellySecurityContextRunAsUser()
		j.SmellySecurityContextCapabilities()
		j.SmellySecurityContextAllowPrivilegeEscalation()
		j.SmellySecurityContextReadOnlyRootFilesystem()
		smells = append(smells, j.SmellJob...)
	}
	return smells
}

func (smellyService SmellyService) FindCronJobSmell(cronJobs []batchv1.CronJob) (smells []models.SmellCronJob) {
	smells = []models.SmellCronJob{}
	for _, cronJob := range cronJobs {
		c := &implementation.CronJob{
			CronJob: &cronJob,
		}
		c.SmellyResourceAndLimit()
		c.SmellySecurityContextRunAsUser()
		c.SmellySecurityContextCapabilities()
		c.SmellySecurityContextAllowPrivilegeEscalation()
		c.SmellySecurityContextReadOnlyRootFilesystem()
		smells = append(smells, c.SmellCronJob...)
	}
	return smells
}

func (smellyService SmellyService) Execute(manifestToFindSmells string) (pods []corev1.Pod, deployments []appsv1.Deployment, statefulsets []appsv1.StatefulSet, daemonsets []appsv1.DaemonSet, jobs []batchv1.Job, cronJobs []batchv1.CronJob, err error) {
	log.Info("Executing smelly service")
	var podSlices []corev1.Pod
	var deploymentSlices []appsv1.Deployment
	var statefulSetSlices []appsv1.StatefulSet
	var daemonSetSlices []appsv1.DaemonSet
	var jobSlices []batchv1.Job
	var cronJobSlices []batchv1.CronJob

	decode := scheme.Codecs.UniversalDeserializer().Decode
	for _, spec := range strings.Split(manifestToFindSmells, "---") {
		if len(spec) == 0 {
			continue
		}
		obj, _, err := decode([]byte(spec), nil, nil)
		if err != nil {
			continue
		}
		switch obj.(type) {
		case *corev1.Pod:
			pods := obj.(*corev1.Pod)
			log.Info("Name:", pods.GetName())
			log.Info("Namespace:", pods.GetNamespace())
			log.Info("Kind:", pods.GetResourceVersion())
			log.Info("---")
			podSlices = append(podSlices, *pods)
		case *appsv1.Deployment:
			d := obj.(*appsv1.Deployment)
			log.Info("Name:", d.GetName())
			log.Info("Namespace:", d.GetNamespace())
			log.Info("GVK:", d.GroupVersionKind())
			log.Info("Containers IMAGEMS", d.Spec.Template.Spec.Containers[0].Image)
			log.Info("---")
			deploymentSlices = append(deploymentSlices, *d)
		case *appsv1.StatefulSet:
			ss := obj.(*appsv1.StatefulSet)
			log.Info("Name:", ss.GetName())
			log.Info("Namespace:", ss.GetNamespace())
			log.Info("GVK:", ss.GroupVersionKind())
			log.Info("---")
			statefulSetSlices = append(statefulSetSlices, *ss)
		case *appsv1.DaemonSet:
			ds := obj.(*appsv1.DaemonSet)
			log.Info("Name:", ds.GetName())
			log.Info("Namespace:", ds.GetNamespace())
			log.Info("GVK:", ds.GroupVersionKind())
			log.Info("---")
			daemonSetSlices = append(daemonSetSlices, *ds)
		case *batchv1.Job:
			job := obj.(*batchv1.Job)
			log.Info("Name:", job.GetName())
			log.Info("Namespace:", job.GetNamespace())
			log.Info("Kind:", job.GetResourceVersion())
			log.Info("---")
			jobSlices = append(jobSlices, *job)
		case *batchv1.CronJob:
			cronJob := obj.(*batchv1.CronJob)
			log.Info("Name:", cronJob.GetName())
			log.Info("Namespace:", cronJob.GetNamespace())
			log.Info("Kind:", cronJob.GetResourceVersion())
			log.Info("---")
			cronJobSlices = append(cronJobSlices, *cronJob)
		}
	}
	if len(podSlices) == 0 && len(deploymentSlices) == 0 && len(statefulSetSlices) == 0 && len(daemonSetSlices) == 0 || len(jobSlices) == 0 || len(cronJobSlices) == 0 {
		log.Info("No pods, deployments, statefulsets or daemonsets found in the manifest")
		return nil, nil, nil, nil, nil, nil, errors.New("no pods, deployments, statefulsets or daemonsets found in the manifest. Please provide a valid manifest with at least one pod, deployment, statefulset or daemonset")
	}
	return podSlices, deploymentSlices, statefulSetSlices, daemonSetSlices, jobSlices, cronJobSlices, nil
}

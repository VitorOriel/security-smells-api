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

func (smellyService SmellyService) FindReplicaSetSmell(replicaSets []appsv1.ReplicaSet) (smells []models.SmellKubernetes) {
	smells = []models.SmellKubernetes{}
	for _, replicaSet := range replicaSets {
		r := &implementation.ReplicaSet{
			ReplicaSet: &replicaSet,
		}
		r.SmellyResourceAndLimit()
		r.SmellySecurityContextRunAsUser()
		r.SmellySecurityContextCapabilities()
		r.SmellySecurityContextAllowPrivilegeEscalation()
		r.SmellySecurityContextReadOnlyRootFilesystem()
		smells = append(smells, r.SmellKubernetes...)
	}
	return smells
}

func (smellyService SmellyService) FindDaemonSetSmell(daemonSets []appsv1.DaemonSet) (smells []models.SmellKubernetes) {
	smells = []models.SmellKubernetes{}
	for _, daemonSet := range daemonSets {
		d := &implementation.DaemonSet{
			DaemonSet: &daemonSet,
		}
		d.SmellyResourceAndLimit()
		d.SmellySecurityContextRunAsUser()
		d.SmellySecurityContextCapabilities()
		d.SmellySecurityContextAllowPrivilegeEscalation()
		d.SmellySecurityContextReadOnlyRootFilesystem()
		smells = append(smells, d.SmellKubernetes...)
	}
	return smells
}

func (smellyService SmellyService) FindStatefulSetSmell(statefulSets []appsv1.StatefulSet) (smells []models.SmellKubernetes) {
	smells = []models.SmellKubernetes{}
	for _, statefulSet := range statefulSets {
		s := &implementation.StatefulSet{
			StatefulSet: &statefulSet,
		}
		s.SmellyResourceAndLimit()
		s.SmellySecurityContextRunAsUser()
		s.SmellySecurityContextCapabilities()
		s.SmellySecurityContextAllowPrivilegeEscalation()
		s.SmellySecurityContextReadOnlyRootFilesystem()
		smells = append(smells, s.SmellKubernetes...)
	}
	return smells
}

func (smellyService SmellyService) FindDeploymentSmell(deployments []appsv1.Deployment) (smells []models.SmellKubernetes) {
	smells = []models.SmellKubernetes{}
	for _, deployment := range deployments {
		d := &implementation.Deployment{
			Deployment: &deployment,
		}
		d.SmellyResourceAndLimit()
		d.SmellySecurityContextRunAsUser()
		d.SmellySecurityContextCapabilities()
		d.SmellySecurityContextAllowPrivilegeEscalation()
		d.SmellySecurityContextReadOnlyRootFilesystem()
		smells = append(smells, d.SmellKubernetes...)
	}

	return smells
}

func (smellyService SmellyService) FindPodSmell(pods []corev1.Pod) (smells []models.SmellKubernetes) {
	smells = []models.SmellKubernetes{}
	for _, pod := range pods {
		p := &implementation.Pod{
			Pod: &pod,
		}
		p.SmellyResourceAndLimit()
		p.SmellySecurityContextRunAsUser()
		p.SmellySecurityContextCapabilities()
		p.SmellySecurityContextAllowPrivilegeEscalation()
		p.SmellySecurityContextReadOnlyRootFilesystem()
		smells = append(smells, p.SmellKubernetes...)
	}
	return smells
}

func (smellyService SmellyService) FindJobSmell(jobs []batchv1.Job) (smells []models.SmellKubernetes) {
	smells = []models.SmellKubernetes{}
	for _, job := range jobs {
		j := &implementation.Job{
			Job: &job,
		}
		j.SmellyResourceAndLimit()
		j.SmellySecurityContextRunAsUser()
		j.SmellySecurityContextCapabilities()
		j.SmellySecurityContextAllowPrivilegeEscalation()
		j.SmellySecurityContextReadOnlyRootFilesystem()
		smells = append(smells, j.SmellKubernetes...)
	}
	return smells
}

func (smellyService SmellyService) FindCronJobSmell(cronJobs []batchv1.CronJob) (smells []models.SmellKubernetes) {
	smells = []models.SmellKubernetes{}
	for _, cronJob := range cronJobs {
		c := &implementation.CronJob{
			CronJob: &cronJob,
		}
		c.SmellyResourceAndLimit()
		c.SmellySecurityContextRunAsUser()
		c.SmellySecurityContextCapabilities()
		c.SmellySecurityContextAllowPrivilegeEscalation()
		c.SmellySecurityContextReadOnlyRootFilesystem()
		smells = append(smells, c.SmellKubernetes...)
	}
	return smells
}

func (smellyService SmellyService) Execute(manifestToFindSmells string) (pods []corev1.Pod, replicaSets []appsv1.ReplicaSet, deployments []appsv1.Deployment, statefulsets []appsv1.StatefulSet, daemonsets []appsv1.DaemonSet, jobs []batchv1.Job, cronJobs []batchv1.CronJob, err error) {
	log.Info("Executing smelly service")
	var podSlices []corev1.Pod
	var replicaSetSlices []appsv1.ReplicaSet
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
		switch obj := obj.(type) {
		case *corev1.Pod:
			log.Info("Name:", obj.GetName())
			log.Info("Namespace:", obj.GetNamespace())
			log.Info("Kind:", obj.GetResourceVersion())
			log.Info("---")
			podSlices = append(podSlices, *obj)
		case *appsv1.ReplicaSet:
			log.Info("Name:", obj.GetName())
			log.Info("Namespace:", obj.GetNamespace())
			log.Info("GVK:", obj.GroupVersionKind())
			log.Info("Containers IMAGEMS", obj.Spec.Template.Spec.Containers[0].Image)
			log.Info("---")
			replicaSetSlices = append(replicaSetSlices, *obj)
		case *appsv1.Deployment:
			log.Info("Name:", obj.GetName())
			log.Info("Namespace:", obj.GetNamespace())
			log.Info("GVK:", obj.GroupVersionKind())
			log.Info("Containers IMAGEMS", obj.Spec.Template.Spec.Containers[0].Image)
			log.Info("---")
			deploymentSlices = append(deploymentSlices, *obj)
		case *appsv1.StatefulSet:
			log.Info("Name:", obj.GetName())
			log.Info("Namespace:", obj.GetNamespace())
			log.Info("GVK:", obj.GroupVersionKind())
			log.Info("---")
			statefulSetSlices = append(statefulSetSlices, *obj)
		case *appsv1.DaemonSet:
			log.Info("Name:", obj.GetName())
			log.Info("Namespace:", obj.GetNamespace())
			log.Info("GVK:", obj.GroupVersionKind())
			log.Info("---")
			daemonSetSlices = append(daemonSetSlices, *obj)
		case *batchv1.Job:
			log.Info("Name:", obj.GetName())
			log.Info("Namespace:", obj.GetNamespace())
			log.Info("Kind:", obj.GetResourceVersion())
			log.Info("---")
			jobSlices = append(jobSlices, *obj)
		case *batchv1.CronJob:
			log.Info("Name:", obj.GetName())
			log.Info("Namespace:", obj.GetNamespace())
			log.Info("Kind:", obj.GetResourceVersion())
			log.Info("---")
			cronJobSlices = append(cronJobSlices, *obj)
		}
	}
	if len(podSlices) == 0 && len(deploymentSlices) == 0 && len(statefulSetSlices) == 0 && len(daemonSetSlices) == 0 && len(jobSlices) == 0 && len(cronJobSlices) == 0 {
		log.Info("No pods, deployments, statefulsets or daemonsets found in the manifest")
		return nil, nil, nil, nil, nil, nil, nil, errors.New("no pods, deployments, statefulsets or daemonsets found in the manifest. Please provide a valid manifest with at least one pod, deployment, statefulset or daemonset")
	}
	return podSlices, replicaSetSlices, deploymentSlices, statefulSetSlices, daemonSetSlices, jobSlices, cronJobSlices, nil
}

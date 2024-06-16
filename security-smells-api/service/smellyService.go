package service

import (
	"errors"
	"security-smells-api/constants"
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

func (smellyService SmellyService) FindReplicaSetSmell(replicaSets []appsv1.ReplicaSet, workloadPosition []int) []*models.SmellKubernetes {
	smells := []*models.SmellKubernetes{}
	for i, replicaSet := range replicaSets {
		r := &implementation.ReplicaSet{
			ReplicaSet:       &replicaSet,
			WorkloadPosition: workloadPosition[i],
		}
		r.SmellyResourceAndLimit()
		r.SmellySecurityContextRunAsUser()
		r.SmellySecurityContextCapabilities()
		r.SmellySecurityContextAllowPrivilegeEscalation()
		r.SmellySecurityContextReadOnlyRootFilesystem()
		r.SmellySecurityContextPrivileged()
		smells = append(smells, r.SmellKubernetes...)
	}
	return smells
}

func (smellyService SmellyService) FindDaemonSetSmell(daemonSets []appsv1.DaemonSet, workloadPosition []int) []*models.SmellKubernetes {
	smells := []*models.SmellKubernetes{}
	for i, daemonSet := range daemonSets {
		d := &implementation.DaemonSet{
			DaemonSet:        &daemonSet,
			WorkloadPosition: workloadPosition[i],
		}
		d.SmellyResourceAndLimit()
		d.SmellySecurityContextRunAsUser()
		d.SmellySecurityContextCapabilities()
		d.SmellySecurityContextAllowPrivilegeEscalation()
		d.SmellySecurityContextReadOnlyRootFilesystem()
		d.SmellySecurityContextPrivileged()
		smells = append(smells, d.SmellKubernetes...)
	}
	return smells
}

func (smellyService SmellyService) FindStatefulSetSmell(statefulSets []appsv1.StatefulSet, workloadPosition []int) []*models.SmellKubernetes {
	smells := []*models.SmellKubernetes{}
	for i, statefulSet := range statefulSets {
		s := &implementation.StatefulSet{
			StatefulSet:      &statefulSet,
			WorkloadPosition: workloadPosition[i],
		}
		s.SmellyResourceAndLimit()
		s.SmellySecurityContextRunAsUser()
		s.SmellySecurityContextCapabilities()
		s.SmellySecurityContextAllowPrivilegeEscalation()
		s.SmellySecurityContextReadOnlyRootFilesystem()
		s.SmellySecurityContextPrivileged()
		smells = append(smells, s.SmellKubernetes...)
	}
	return smells
}

func (smellyService SmellyService) FindDeploymentSmell(deployments []appsv1.Deployment, workloadPosition []int) []*models.SmellKubernetes {
	smells := []*models.SmellKubernetes{}
	for i, deployment := range deployments {
		d := &implementation.Deployment{
			Deployment:       &deployment,
			WorkloadPosition: workloadPosition[i],
		}
		d.SmellyResourceAndLimit()
		d.SmellySecurityContextRunAsUser()
		d.SmellySecurityContextCapabilities()
		d.SmellySecurityContextAllowPrivilegeEscalation()
		d.SmellySecurityContextReadOnlyRootFilesystem()
		d.SmellySecurityContextPrivileged()
		smells = append(smells, d.SmellKubernetes...)
	}
	return smells
}

func (smellyService SmellyService) FindPodSmell(pods []corev1.Pod, workloadPosition []int) []*models.SmellKubernetes {
	smells := []*models.SmellKubernetes{}
	for i, pod := range pods {
		p := &implementation.Pod{
			Pod:              &pod,
			WorkloadPosition: workloadPosition[i],
		}
		p.SmellyResourceAndLimit()
		p.SmellySecurityContextRunAsUser()
		p.SmellySecurityContextCapabilities()
		p.SmellySecurityContextAllowPrivilegeEscalation()
		p.SmellySecurityContextReadOnlyRootFilesystem()
		p.SmellySecurityContextPrivileged()
		smells = append(smells, p.SmellKubernetes...)
	}
	return smells
}

func (smellyService SmellyService) FindJobSmell(jobs []batchv1.Job, workloadPosition []int) []*models.SmellKubernetes {
	smells := []*models.SmellKubernetes{}
	for i, job := range jobs {
		j := &implementation.Job{
			Job:              &job,
			WorkloadPosition: workloadPosition[i],
		}
		j.SmellyResourceAndLimit()
		j.SmellySecurityContextRunAsUser()
		j.SmellySecurityContextCapabilities()
		j.SmellySecurityContextAllowPrivilegeEscalation()
		j.SmellySecurityContextReadOnlyRootFilesystem()
		j.SmellySecurityContextPrivileged()
		smells = append(smells, j.SmellKubernetes...)
	}
	return smells
}

func (smellyService SmellyService) FindCronJobSmell(cronJobs []batchv1.CronJob, workloadPosition []int) []*models.SmellKubernetes {
	smells := []*models.SmellKubernetes{}
	for i, cronJob := range cronJobs {
		c := &implementation.CronJob{
			CronJob:          &cronJob,
			WorkloadPosition: workloadPosition[i],
		}
		c.SmellyResourceAndLimit()
		c.SmellySecurityContextRunAsUser()
		c.SmellySecurityContextCapabilities()
		c.SmellySecurityContextAllowPrivilegeEscalation()
		c.SmellySecurityContextReadOnlyRootFilesystem()
		c.SmellySecurityContextPrivileged()
		smells = append(smells, c.SmellKubernetes...)
	}
	return smells
}

func (smellyService SmellyService) Execute(manifestToFindSmells string) (*models.KubernetesWorkloads, *models.FileMetadata[constants.KubernetesWorkload], error) {
	log.Info("Executing smelly service")
	kubernetesWorkloads := new(models.KubernetesWorkloads)
	fileMetadata := models.NewFileMetadata[constants.KubernetesWorkload]()
	decode := scheme.Codecs.UniversalDeserializer().Decode
	for i, spec := range strings.Split(manifestToFindSmells, "---") {
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
			fileMetadata.AppendWorkloadPosition(constants.POD_WORKLOAD, i)
			kubernetesWorkloads.Pods = append(kubernetesWorkloads.Pods, *obj)
		case *appsv1.ReplicaSet:
			log.Info("Name:", obj.GetName())
			log.Info("Namespace:", obj.GetNamespace())
			log.Info("GVK:", obj.GroupVersionKind())
			log.Info("---")
			fileMetadata.AppendWorkloadPosition(constants.REPLICASET_WORKLOAD, i)
			kubernetesWorkloads.ReplicaSets = append(kubernetesWorkloads.ReplicaSets, *obj)
		case *appsv1.Deployment:
			log.Info("Name:", obj.GetName())
			log.Info("Namespace:", obj.GetNamespace())
			log.Info("GVK:", obj.GroupVersionKind())
			log.Info("---")
			fileMetadata.AppendWorkloadPosition(constants.DEPLOYMENT_WORKLOAD, i)
			kubernetesWorkloads.Deployments = append(kubernetesWorkloads.Deployments, *obj)
		case *appsv1.StatefulSet:
			log.Info("Name:", obj.GetName())
			log.Info("Namespace:", obj.GetNamespace())
			log.Info("GVK:", obj.GroupVersionKind())
			log.Info("---")
			fileMetadata.AppendWorkloadPosition(constants.STATEFULSET_WORKLOAD, i)
			kubernetesWorkloads.StatefulSets = append(kubernetesWorkloads.StatefulSets, *obj)
		case *appsv1.DaemonSet:
			log.Info("Name:", obj.GetName())
			log.Info("Namespace:", obj.GetNamespace())
			log.Info("GVK:", obj.GroupVersionKind())
			log.Info("---")
			fileMetadata.AppendWorkloadPosition(constants.DAEMONSET_WORKLOAD, i)
			kubernetesWorkloads.DaemonSets = append(kubernetesWorkloads.DaemonSets, *obj)
		case *batchv1.Job:
			log.Info("Name:", obj.GetName())
			log.Info("Namespace:", obj.GetNamespace())
			log.Info("Kind:", obj.GetResourceVersion())
			log.Info("---")
			fileMetadata.AppendWorkloadPosition(constants.JOB_WORKLOAD, i)
			kubernetesWorkloads.Jobs = append(kubernetesWorkloads.Jobs, *obj)
		case *batchv1.CronJob:
			log.Info("Name:", obj.GetName())
			log.Info("Namespace:", obj.GetNamespace())
			log.Info("Kind:", obj.GetResourceVersion())
			log.Info("---")
			fileMetadata.AppendWorkloadPosition(constants.CRONJOB_WORKLOAD, i)
			kubernetesWorkloads.CronJobs = append(kubernetesWorkloads.CronJobs, *obj)
		}
	}
	if kubernetesWorkloads.IsEmpty() {
		log.Info("could not load any kubernetes workload from provided resource")
		return nil, nil, errors.New("could not load any kubernetes workload from provided resource")
	}
	return kubernetesWorkloads, fileMetadata, nil
}

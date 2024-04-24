package service

import (
	"errors"
	"github.com/gofiber/fiber/v2/log"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"security-smells-api/models"
	"security-smells-api/repository"
	"security-smells-api/service/implementation"
	"strings"
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
		//smells = append(smells, d.SmellDeployment...)
	}

	return smells
}

func (smellyService SmellyService) Execute(manifestToFindSmells string) (pods []corev1.Pod, deployments []appsv1.Deployment, statefulsets []appsv1.StatefulSet, daemonsets []appsv1.DaemonSet, err error) {
	log.Info("Executing smelly service")
	var podSlices []corev1.Pod
	var deploymentSlices []appsv1.Deployment
	var statefulSetSlices []appsv1.StatefulSet
	var daemonSetSlices []appsv1.DaemonSet

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
		}
	}
	if len(podSlices) == 0 && len(deploymentSlices) == 0 && len(statefulSetSlices) == 0 && len(daemonSetSlices) == 0 {
		log.Info("No pods, deployments, statefulsets or daemonsets found in the manifest")
		return nil, nil, nil, nil, errors.New("no pods, deployments, statefulsets or daemonsets found in the manifest. Please provide a valid manifest with at least one pod, deployment, statefulset or daemonset")
	}
	return podSlices, deploymentSlices, statefulSetSlices, daemonSetSlices, nil
}

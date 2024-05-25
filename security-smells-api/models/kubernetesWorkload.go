package models

import (
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
)

type KubernetesWorkloads struct {
	Pods         []corev1.Pod
	ReplicaSets  []appsv1.ReplicaSet
	Deployments  []appsv1.Deployment
	StatefulSets []appsv1.StatefulSet
	DaemonSets   []appsv1.DaemonSet
	Jobs         []batchv1.Job
	CronJobs     []batchv1.CronJob
}

func (k *KubernetesWorkloads) IsEmpty() bool {
	return len(k.Pods) == 0 && len(k.ReplicaSets) == 0 && len(k.Deployments) == 0 && len(k.StatefulSets) == 0 && len(k.DaemonSets) == 0 && len(k.Jobs) == 0 && len(k.CronJobs) == 0
}

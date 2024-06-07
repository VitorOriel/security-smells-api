package constants

type KubernetesWorkload string

const (
	POD_WORKLOAD         KubernetesWorkload = "POD_WORKLOAD"
	JOB_WORKLOAD         KubernetesWorkload = "JOB_WORKLOAD"
	CRONJOB_WORKLOAD     KubernetesWorkload = "CRONJOB_WORKLOAD"
	REPLICASET_WORKLOAD  KubernetesWorkload = "REPLICASET_WORKLOAD"
	DEPLOYMENT_WORKLOAD  KubernetesWorkload = "DEPLOYMENT_WORKLOAD"
	DAEMONSET_WORKLOAD   KubernetesWorkload = "DAEMONSET_WORKLOAD"
	STATEFULSET_WORKLOAD KubernetesWorkload = "STATEFULSET_WORKLOAD"
)

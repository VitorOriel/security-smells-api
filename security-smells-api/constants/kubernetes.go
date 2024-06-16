package constants

type KubernetesWorkload string
type KubernetesRule string

const (
	POD_WORKLOAD         KubernetesWorkload = "POD_WORKLOAD"
	JOB_WORKLOAD         KubernetesWorkload = "JOB_WORKLOAD"
	CRONJOB_WORKLOAD     KubernetesWorkload = "CRONJOB_WORKLOAD"
	REPLICASET_WORKLOAD  KubernetesWorkload = "REPLICASET_WORKLOAD"
	DEPLOYMENT_WORKLOAD  KubernetesWorkload = "DEPLOYMENT_WORKLOAD"
	DAEMONSET_WORKLOAD   KubernetesWorkload = "DAEMONSET_WORKLOAD"
	STATEFULSET_WORKLOAD KubernetesWorkload = "STATEFULSET_WORKLOAD"
)

const (
	K8S_SEC_CAPABILITIES   KubernetesRule = "K8S_SEC_CAPABILITIES"
	K8S_SEC_RUNASUSER      KubernetesRule = "K8S_SEC_RUNASUSER"
	K8S_SEC_PRIVESCALATION KubernetesRule = "K8S_SEC_PRIVESCALATION"
	K8S_SEC_ROROOTFS       KubernetesRule = "K8S_SEC_ROROOTFS"
	K8S_SEC_RESREQUESTS    KubernetesRule = "K8S_SEC_RESREQUESTS"
	K8S_SEC_RESLIMITS      KubernetesRule = "K8S_SEC_RESLIMITS"
	K8S_SEC_PRIVILEGED     KubernetesRule = "K8S_SEC_PRIVILEGED"
)

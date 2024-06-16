package models

type Meta struct {
	TotalOfSmells int           `json:"totalOfSmells"`
	Workload      *workloadMeta `json:"workload"`
}

type Data struct {
	SmellsReplicaSet  []*SmellKubernetes `json:"smellsReplicaSet"`
	SmellsDeployment  []*SmellKubernetes `json:"smellsDeployment"`
	SmellsPod         []*SmellKubernetes `json:"smellsPod"`
	SmellsJob         []*SmellKubernetes `json:"smellsJob"`
	SmellsCronJob     []*SmellKubernetes `json:"smellsCronJob"`
	SmellsStatefulSet []*SmellKubernetes `json:"smellsStatefulSet"`
	SmellsDaemonSet   []*SmellKubernetes `json:"smellsDaemonSet"`
}

type SmellyResponseDTO struct {
	Meta *Meta `json:"meta"`
	Data *Data `json:"data"`
}

type workloadMeta struct {
	DecodedReplicaSets  int `json:"decodedReplicaSets"`
	DecodedDeployments  int `json:"decodedDeployments"`
	DecodedPods         int `json:"decodedPods"`
	DecodedJobs         int `json:"decodedJobs"`
	DecodedCronJobs     int `json:"decodedCronJobs"`
	DecodedStatefulSets int `json:"decodedStatefulSets"`
	DecodedDaemonSets   int `json:"decodedDaemonSets"`
}

func NewWorkloadMeta(kubernetesWorkloads *KubernetesWorkloads) *workloadMeta {
	return &workloadMeta{
		DecodedReplicaSets:  len(kubernetesWorkloads.ReplicaSets),
		DecodedDeployments:  len(kubernetesWorkloads.Deployments),
		DecodedPods:         len(kubernetesWorkloads.Pods),
		DecodedJobs:         len(kubernetesWorkloads.Jobs),
		DecodedCronJobs:     len(kubernetesWorkloads.CronJobs),
		DecodedStatefulSets: len(kubernetesWorkloads.StatefulSets),
		DecodedDaemonSets:   len(kubernetesWorkloads.DaemonSets),
	}
}

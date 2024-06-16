package models

type Meta struct {
	TotalOfSmells    int                  `json:"totalOfSmells"`
	DecodedWorkloads *decodedWorkloadMeta `json:"decodedWorkloads"`
}

type Data struct {
	SmellsReplicaSet  []*SmellKubernetes `json:"ReplicaSet"`
	SmellsDeployment  []*SmellKubernetes `json:"Deployment"`
	SmellsPod         []*SmellKubernetes `json:"Pod"`
	SmellsJob         []*SmellKubernetes `json:"Job"`
	SmellsCronJob     []*SmellKubernetes `json:"CronJob"`
	SmellsStatefulSet []*SmellKubernetes `json:"StatefulSet"`
	SmellsDaemonSet   []*SmellKubernetes `json:"DaemonSet"`
}

type SmellyResponseDTO struct {
	Meta *Meta `json:"meta"`
	Data *Data `json:"data"`
}

type decodedWorkloadMeta struct {
	ReplicaSets  int `json:"ReplicaSets"`
	Deployments  int `json:"Deployments"`
	Pods         int `json:"Pods"`
	Jobs         int `json:"Jobs"`
	CronJobs     int `json:"CronJobs"`
	StatefulSets int `json:"StatefulSets"`
	DaemonSets   int `json:"DaemonSets"`
}

func NewWorkloadMeta(kubernetesWorkloads *KubernetesWorkloads) *decodedWorkloadMeta {
	return &decodedWorkloadMeta{
		ReplicaSets:  len(kubernetesWorkloads.ReplicaSets),
		Deployments:  len(kubernetesWorkloads.Deployments),
		Pods:         len(kubernetesWorkloads.Pods),
		Jobs:         len(kubernetesWorkloads.Jobs),
		CronJobs:     len(kubernetesWorkloads.CronJobs),
		StatefulSets: len(kubernetesWorkloads.StatefulSets),
		DaemonSets:   len(kubernetesWorkloads.DaemonSets),
	}
}

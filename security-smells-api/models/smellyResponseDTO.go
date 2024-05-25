package models

type Meta struct {
	TotalOfSmells int `json:"totalOfSmells"`
}

type Data struct {
	SmellsReplicaSet  []SmellReplicaSet  `json:"smellsReplicaSet"`
	SmellsDeployment  []SmellDeployment  `json:"smellsDeployment"`
	SmellsPod         []SmellPod         `json:"smellsPod"`
	SmellsJob         []SmellJob         `json:"smellsJob"`
	SmellsCronJob     []SmellCronJob     `json:"smellsCronJob"`
	SmellsStatefulSet []SmellStatefulSet `json:"smellsStatefulSet"`
	SmellDemonSet     []SmellDaemonSet   `json:"smellDemonSet"`
}

type SmellyResponseDTO struct {
	Meta Meta `json:"meta"`
	Data Data `json:"data"`
}

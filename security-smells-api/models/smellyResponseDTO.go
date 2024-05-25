package models

type Meta struct {
	TotalOfSmells int `json:"totalOfSmells"`
}

type Data struct {
	SmellsReplicaSet  []SmellKubernetes `json:"smellsReplicaSet"`
	SmellsDeployment  []SmellKubernetes `json:"smellsDeployment"`
	SmellsPod         []SmellKubernetes `json:"smellsPod"`
	SmellsJob         []SmellKubernetes `json:"smellsJob"`
	SmellsCronJob     []SmellKubernetes `json:"smellsCronJob"`
	SmellsStatefulSet []SmellKubernetes `json:"smellsStatefulSet"`
	SmellDemonSet     []SmellKubernetes `json:"smellDemonSet"`
}

type SmellyResponseDTO struct {
	Meta Meta `json:"meta"`
	Data Data `json:"data"`
}

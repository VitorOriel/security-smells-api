package models

type SmellyResponseDTO struct {
	TotalOfSmells    int               `json:"totalOfSmells"`
	SmellsReplicaSet []SmellReplicaSet `json:"smellsReplicaSet"`
	SmellsDeployment []SmellDeployment `json:"smellsDeployment"`
	SmellsPod        []SmellPod        `json:"smellsPod"`
	SmellsJob        []SmellJob        `json:"smellsJob"`
	SmellsCronJob    []SmellCronJob    `json:"smellsCronJob"`
}

package models

type SmellyResponseDTO struct {
	TotalOfSmells    int               `json:"totalOfSmells"`
	SmellsDeployment []SmellDeployment `json:"smellsDeployment"`
	SmellsPod        []SmellPod        `json:"smellsPod"`
	SmellsJob        []SmellJob        `json:"smellsJob"`
	SmellsCronJob    []SmellCronJob    `json:"smellsCronJob"`
}

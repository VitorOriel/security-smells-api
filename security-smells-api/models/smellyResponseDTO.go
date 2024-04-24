package models

type SmellyResponseDTO struct {
	TotalOfSmells    int               `json:"totalOfSmells"`
	SmellsDeployment []SmellDeployment `json:"smellsDeployment"`
	SmellsPod        []SmellPod        `json:"smellsPod"`
}

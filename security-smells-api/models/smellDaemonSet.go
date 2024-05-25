package models

type SmellDaemonSet struct {
	FileName       string `json:"fileName" validate:"required"`
	NameSpace      string `json:"namespace" validate:"required"`
	DemonSetName   string `json:"daemonSetName" validate:"required"`
	ContainerName  string `json:"containerName" validate:"required"`
	ContainerImage string `json:"containerImage" validate:"required"`
	Kind           string `json:"kind" validate:"required"`
	Message        string `json:"message" validate:"required"`
	Suggestion     string `json:"suggestion" validate:"required"`
}

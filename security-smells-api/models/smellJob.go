package models

type SmellJob struct {
	FileName       string `json:"fileName" validate:"required"`
	NameSpace      string `json:"namespace" validate:"required"`
	JobName        string `json:"jobName" validate:"required"`
	ContainerName  string `json:"containerName" validate:"required"`
	ContainerImage string `json:"containerImage" validate:"required"`
	Kind           string `json:"kind" validate:"required"`
	Message        string `json:"message" validate:"required"`
	Suggestion     string `json:"suggestion" validate:"required"`
}

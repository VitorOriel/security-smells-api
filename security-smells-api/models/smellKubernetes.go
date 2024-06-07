package models

type SmellKubernetes struct {
	Filename          string `json:"filename" validate:"required"`
	Namespace         string `json:"namespace" validate:"required"`
	WorkloadKind      string `json:"workload_kind" validate:"required"`
	WorkloadLabelName string `json:"workload_label_name" validate:"required"`
	WorkloadPosition  int    `json:"workload_position" validate:"required"`
	ContainerName     string `json:"container_name" validate:"required"`
	ContainerImage    string `json:"container_image" validate:"required"`
	Message           string `json:"message" validate:"required"`
	Suggestion        string `json:"suggestion" validate:"required"`
}

package models

type SmellyResponseErrorDTO struct {
	YamlToValidate string `json:"yamlToValidate"`
	Message        string `json:"message"`
}

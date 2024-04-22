package models

type Smelly struct {
	FileName       string `json:"fileName" validate:"required"`
	YamlToValidate string `yaml:"yamlToValidate" validate:"required"`
}

func (smelly Smelly) String() string {
	return "Smelly{" +
		"FileName: " + smelly.FileName + ", " +
		"YamlToValidate: " + smelly.YamlToValidate + ", " +
		"}"
}

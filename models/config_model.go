package models

type AppConfig struct {
	DataPath string `json:"dataPath"`
	UseCache bool   `json:"useCache"`
}

package config

type Config struct {
	DBHost    string `json:"dbHost"`
	DBPort    int    `json:"dbPort"`
	Endpoints []struct {
		Url      string `json:"url"`
		AuthType string `json:"authType"`
	} `json:"endpoints"`
}

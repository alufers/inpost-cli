package main

import "github.com/alufers/inpost-cli/swagger"

type ConfigAccount struct {
	PhoneNumber  string `json:"phoneNumber"`
	AuthToken    string `json:"authToken"`
	RefreshToken string `json:"refreshToken"`
}

func (c *ConfigAccount) ToClient() *swagger.APIClient {
	cfg := swagger.NewConfiguration()
	cfg.DefaultHeader["Authorization"] = "Bearer " + c.AuthToken
	apiClient := swagger.NewAPIClient(cfg)
	return apiClient
}

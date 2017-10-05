package tests

import (
	"encoding/json"
	"fmt"
	"os"
)

type envParams struct {
	AccountWebAddress string `json:"accountWebAddress"`
	UserHash          string `json:"userHash"`
	UserLogin         string `json:"userLogin"`
}

func newEnvParams() (*envParams, error) {
	p := &envParams{
		AccountWebAddress: os.Getenv("AMOCRM_ACOOUNT_WEB_ADDRESS"),
		UserHash:          os.Getenv("AMOCRM_USER_HASH"),
		UserLogin:         os.Getenv("AMOCRM_USER_LOGIN"),
	}

	if len(p.AccountWebAddress) > 0 && len(p.UserHash) > 0 && len(p.UserLogin) > 0 {
		return p, nil
	} else if len(p.AccountWebAddress) > 0 || len(p.UserHash) > 0 || len(p.UserLogin) > 0 {
		return nil, fmt.Errorf("You need to set all environmet variables AMOCRM_ACOOUNT_WEB_ADDRESS and AMOCRM_USER_HASH and AMOCRM_USER_LOGIN")
	}

	envFile, err := os.Open("../env.json")
	if err != nil {
		return nil, err
	}
	defer envFile.Close()
	err = json.NewDecoder(envFile).Decode(p)
	return p, err
}

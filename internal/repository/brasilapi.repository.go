package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alextavella/multithreading/internal/provider"
)

type BrasilAPIRepository struct {
	Name string
}

func NewBrasilAPIRepository() provider.IAddressProvider {
	return &BrasilAPIRepository{
		Name: "BrasilAPI",
	}
}

func (r *BrasilAPIRepository) ProviderName() string {
	return r.Name
}

func (r *BrasilAPIRepository) SearchByZipCode(zipcode string) (*provider.SearchAddressByZipCodeResult, error) {
	req, _ := http.NewRequest(http.MethodGet, "https://brasilapi.com.br/api/cep/v1/"+zipcode, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, provider.ErrSearchAddressByZipCode
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &provider.SearchAddressByZipCodeResult{
			Address: "",
		}, nil
	}

	var apiResult BrasilAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResult)
	if err != nil {
		return nil, provider.ErrSearchAddressByZipCode
	}

	var result = provider.SearchAddressByZipCodeResult{
		Address: fmt.Sprintf("%s, %s, %s - %s", apiResult.Street, apiResult.Neighborhood, apiResult.City, apiResult.State),
	}

	return &result, nil
}

/**
API Response:
{
	"cep": "09861160",
	"state": "SP",
	"city": "São Bernardo do Campo",
	"neighborhood": "Independência",
	"street": "Avenida Moinho Fabrini",
	"service": "open-cep"
}
**/

type BrasilAPIResponse struct {
	CEP          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

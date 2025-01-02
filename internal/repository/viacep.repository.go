package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alextavella/multithreading/internal/provider"
)

type ViaCEPRepository struct {
	Name string
}

func NewViaCEPRepository() provider.IAddressProvider {
	return &ViaCEPRepository{
		Name: "ViaCEP",
	}
}

func (r *ViaCEPRepository) ProviderName() string {
	return r.Name
}

func (r *ViaCEPRepository) SearchByZipCode(zipcode string) (*provider.SearchAddressByZipCodeResult, error) {
	req, _ := http.NewRequest(http.MethodGet, "http://viacep.com.br/ws/"+zipcode+"/json/", nil)
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

	var apiResult ViaCEPResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResult)
	if err != nil {
		return nil, provider.ErrSearchAddressByZipCode
	}

	var result = provider.SearchAddressByZipCodeResult{
		Address: fmt.Sprintf("%s, %s, %s - %s", apiResult.Logradouro, apiResult.Bairro, apiResult.Localidade, apiResult.Estado),
	}

	return &result, nil
}

/**
API Response:
{
	"cep": "09861-160",
	"logradouro": "Avenida Moinho Fabrini",
	"complemento": "até 609/610",
	"unidade": "",
	"bairro": "Independência",
	"localidade": "São Bernardo do Campo",
	"uf": "SP",
	"estado": "São Paulo",
	"regiao": "Sudeste",
	"ibge": "3548708",
	"gia": "6350",
	"ddd": "11",
	"siafi": "7075"
}
**/

type ViaCEPResponse struct {
	CEP        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Estado     string `json:"estado"`
	Localidade string `json:"localidade"`
	Bairro     string `json:"bairro"`
}

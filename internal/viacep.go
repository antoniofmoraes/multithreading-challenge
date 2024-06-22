package internal

import (
	"encoding/json"
	"fmt"
)

type ViacepCepResponse struct {
	Cep        string `json:"cep"`
	Uf         string `json:"uf"`
	Cidade     string `json:"localidade"`
	Bairro     string `json:"bairro"`
	Logradouro string `json:"logradouro"`
}

func (res *ViacepCepResponse) ToCepInfoResponse() *CepInfoResponse {
	return &CepInfoResponse{
		Cep:          RemoveNonDigits(res.Cep),
		State:        res.Uf,
		City:         res.Cidade,
		Neighborhood: res.Bairro,
		Street:       res.Logradouro,
	}
}

func GetCepInfoFromViacep(cep string) (*CepInfoResponse, error) {
	bytes, err := Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json", cep))
	if err != nil {
		return nil, err
	}

	response := &ViacepCepResponse{}
	err = json.Unmarshal(bytes, response)
	if err != nil {
		return nil, err
	}

	return response.ToCepInfoResponse(), nil
}

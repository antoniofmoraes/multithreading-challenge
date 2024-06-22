package internal

import (
	"encoding/json"
	"fmt"
)

func GetCepInfoFromBrasilapi(cep string) (*CepInfoResponse, error) {
	bytes, err := Get(fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep))
	if err != nil {
		return nil, err
	}

	cepInfoResponse := &CepInfoResponse{}
	err = json.Unmarshal(bytes, cepInfoResponse)
	if err != nil {
		return nil, err
	}

	return cepInfoResponse, nil
}

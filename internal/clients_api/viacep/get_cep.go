package viacep

import (
	"fmt"

	clientsApi "github.com/natashaCarreao/go-multithreading/internal/clients_api"
	"github.com/natashaCarreao/go-multithreading/internal/clients_api/responses"
)

const (
	urlBase = "https://viacep.com.br/ws/%s/json/"
)

func (c *client) GetCEP(cep string) (*responses.ViaCep, error) {

	url := fmt.Sprintf(urlBase, cep)

	reader, err := clientsApi.Get(url, 1)
	if err != nil {
		return &responses.ViaCep{}, err
	}

	var cepResponse *responses.ViaCep
	err = clientsApi.Decode(reader, &cepResponse)
	if err != nil {
		return &responses.ViaCep{}, err
	}

	return cepResponse, nil

}

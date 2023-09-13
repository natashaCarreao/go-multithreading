package viacep

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/natashaCarreao/go-multithreading/internal/clients_api/responses"
	"github.com/natashaCarreao/go-multithreading/internal/domains"
)

const (
	urlBase = "https://viacep.com.br/ws/%s/json/"
)

func (c client) GetCEP(cep string) (*domains.Cep, error) {

	url := fmt.Sprintf(urlBase, cep)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return &domains.Cep{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &domains.Cep{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &domains.Cep{}, err
	}
	var cepResponse responses.ViaCep
	err = json.Unmarshal(body, &cepResponse)
	if err != nil {
		return &domains.Cep{}, err
	}

	return &domains.Cep{
		Cep:        cepResponse.Cep,
		Logradouro: cepResponse.Logradouro,
		Bairro:     cepResponse.Bairro,
		Cidade:     cepResponse.Localidade,
		Uf:         cepResponse.Uf,
	}, nil

}

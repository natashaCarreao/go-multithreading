package cdn

import (
	"fmt"

	clientsApi "github.com/natashaCarreao/go-multithreading/internal/clients_api"
	"github.com/natashaCarreao/go-multithreading/internal/clients_api/responses"
)

const (
	urlBase = "https://cdn.apicep.com/file/apicep/%s.json"
)

func (c *client) GetCEP(cep string) (*responses.CdnCep, error) {

	url := fmt.Sprintf(urlBase, cep)

	reader, err := clientsApi.Get(url, 1)
	if err != nil {
		return &responses.CdnCep{}, err
	}

	var cepResponse *responses.CdnCep
	err = clientsApi.Decode(reader, &cepResponse)
	if err != nil {
		return cepResponse, err
	}

	return cepResponse, nil
}

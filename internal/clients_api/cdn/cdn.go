package cdn

import (
	"github.com/natashaCarreao/go-multithreading/internal/clients_api/responses"
)

type client struct {
}
type Client interface {
	GetCEP(cep string) (*responses.CdnCep, error)
}

func NewClient() Client {
	return &client{}
}

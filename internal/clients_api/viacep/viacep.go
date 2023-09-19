package viacep

import (
	"github.com/natashaCarreao/go-multithreading/internal/clients_api/responses"
)

type client struct {
}

type Client interface {
	GetCEP(cep string) (*responses.ViaCep, error)
}

func NewClient() Client {
	return &client{}
}

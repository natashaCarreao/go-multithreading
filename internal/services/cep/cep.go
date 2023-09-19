package cep

import (
	"github.com/natashaCarreao/go-multithreading/internal/clients_api/cdn"
	"github.com/natashaCarreao/go-multithreading/internal/clients_api/viacep"
	"github.com/natashaCarreao/go-multithreading/internal/domains"
)

type service struct {
	cdnApi cdn.Client
	viaCep viacep.Client
}

type Service interface {
	GetCEP(cep string) (*domains.Cep, error)
}

func NewService() Service {
	return &service{
		cdnApi: cdn.NewClient(),
		viaCep: viacep.NewClient(),
	}
}

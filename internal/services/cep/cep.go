package cep

import (
	"github.com/natashaCarreao/go-multithreading/internal/clients_api/cdn"
	"github.com/natashaCarreao/go-multithreading/internal/clients_api/viacep"
)

type service struct {
	cdnApi cdn.Client
	viaCep viacep.Client
}

type Service interface {
	GetCEP(ceps []string)
}

func NewService() Service {
	return &service{
		cdnApi: cdn.NewClient(),
		viaCep: viacep.NewClient(),
	}
}

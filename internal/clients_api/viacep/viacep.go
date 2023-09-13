package viacep

import "github.com/natashaCarreao/go-multithreading/internal/domains"

type client struct {
}

type Client interface {
	GetCEP(cep string) (*domains.Cep, error)
}

func NewClient() Client {
	return &client{}
}

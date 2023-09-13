package viacep

type client struct {
}

type Client interface {
	GetCEP(cep string) (string, error)
}

func NewClient() Client {
	return &client{}
}

package services

type service struct {
}

type Service interface {
}

func NewService() Service {
	return &service{}
}

package cep

import (
	"fmt"
	"log"

	"github.com/natashaCarreao/go-multithreading/internal/domains"
)

func (s *service) GetCEP(cep string) (*domains.Cep, error) {

	var cepRespChan = make(chan *domains.Cep)
	var err error

	go s.workerViaCep(cep, cepRespChan)
	go s.workerCdn(cep, cepRespChan)

	tst := <-cepRespChan
	log.Printf("response: %v", tst)
	return nil, err
}

func (s *service) workerViaCep(cep string, cepResp chan *domains.Cep) {
	println(fmt.Sprintf(" ---------- VIACEP:"))
	viacepResp, err := s.viaCep.GetCEP(cep)
	if err != nil {
		log.Printf("Error in find CEP: %s in viacep api: %v", cep, err)
		return
	}
	cepResp <- &domains.Cep{
		Cep:        viacepResp.Cep,
		Logradouro: viacepResp.Logradouro,
		Bairro:     viacepResp.Bairro,
		Cidade:     viacepResp.Localidade,
		Uf:         viacepResp.Uf,
		OrigenApi:  "viacep",
	}
}

func (s *service) workerCdn(cep string, cepResp chan *domains.Cep) {
	println(fmt.Sprintf("************ CDN:"))
	cdnResp, err := s.cdnApi.GetCEP(cep)
	if err != nil {
		log.Printf("Error in find CEP: %s in cdn api: %v", cep, err)
		return
	}
	cepResp <- &domains.Cep{
		Cep:        cdnResp.Code,
		Logradouro: cdnResp.Address,
		Bairro:     cdnResp.District,
		Cidade:     cdnResp.City,
		Uf:         cdnResp.State,
		OrigenApi:  "cdn",
	}
}

package cep

import (
	"log"
	"time"

	"github.com/natashaCarreao/go-multithreading/internal/domains"
)

func (s *service) GetCEP(ceps []string) {

	var viaCepChan = make(chan *domains.Cep)
	var cdnChan = make(chan *domains.Cep)

	for _, cep := range ceps {
		go s.workerViaCep(cep, viaCepChan)
		go s.workerCdn(cep, cdnChan)
	}

	for {
		select {
		case cepResp := <-viaCepChan:
			log.Printf("Found CEP: %s in VIACEP API: [%v]", cepResp.Cep, cepResp)

		case cepResp := <-cdnChan:
			log.Printf("Found CEP: %s in CDN API: [%v]", cepResp.Cep, cepResp)

		case <-time.After(5 * time.Second):
			return
		}

	}
}

func (s *service) workerViaCep(cep string, cepResp chan<- *domains.Cep) {
	viacepResp, err := s.viaCep.GetCEP(cep)
	if err != nil {
		log.Printf("Error in find CEP: %s in viacep api: %v", cep, err)
		return
	}

	if viacepResp.Cep == "" {
		log.Printf("No address found for cep %s", cep)
		return
	}
	cepResp <- &domains.Cep{
		Cep:        viacepResp.Cep,
		Logradouro: viacepResp.Logradouro,
		Bairro:     viacepResp.Bairro,
		Cidade:     viacepResp.Localidade,
		Uf:         viacepResp.Uf,
	}

}

func (s *service) workerCdn(cep string, cepResp chan<- *domains.Cep) {
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
	}
}

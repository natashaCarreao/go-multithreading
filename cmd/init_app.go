package cmd

import (
	"log"

	"github.com/natashaCarreao/go-multithreading/internal/services/cep"
)

func Initialize() error {

	ceps := []string{
		"06180-010",
		"12042-230",
		"12042-231",
		"12042-232",
		"12042-233",
		"12042-234",
		"12042-235",
		"12042-236",
		"12042-237",
		"12042-238",
		"12042-239",
		"12042-240",
	}

	service := cep.NewService()
	service.GetCEP(ceps)
	log.Printf("Finished!")

	return nil
}

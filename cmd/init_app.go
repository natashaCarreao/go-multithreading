package cmd

import (
	"github.com/natashaCarreao/go-multithreading/internal/services/cep"
)

func Initialize() error {

	for i := 0; i < 10; i++ {

		tst := cep.NewService()
		tst.GetCEP("12042-230")

	}

	return nil
}

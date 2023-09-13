package cmd

import "net/http"

func Initialize() error {

	//http.HandleFunc("/", views.GetCotacao)

	return http.ListenAndServe(":8080", nil)
}

package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI               string
	Metodo            string
	Funcao            func(http.ResponseWriter, *http.Request)
	RequerAutorizacao bool
}

func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	rotas = append(rotas, rotasUsuarios...)
	rotas = append(rotas, rotaPaginaPrincial)
	for _, rota := range rotas {
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	fileserver := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets").Handler(http.StripPrefix("/assets/", fileserver))

	return router
}

package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*
			if _, erro := cookies.Ler(r); erro != nil {
				http.Redirect(w, r, "/login", http.StatusMovedPermanently)
				return
			}
		*/
		fmt.Println("autenticando")
		proximaFuncao(w, r)
	}
}
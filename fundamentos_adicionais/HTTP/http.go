package main

import (
	"log"
	"net/http"
)

func raiz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("raiz"))
}

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá mundo"))
	})
	http.HandleFunc("/", raiz)

	log.Fatal(http.ListenAndServe(":5000", nil))
}

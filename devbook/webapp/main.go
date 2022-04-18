package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/routers"
	"webapp/src/utils"
)

func main() {
	utils.CarregarTemplates()
	r := routers.Gerar()

	fmt.Println("Executando webapp na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

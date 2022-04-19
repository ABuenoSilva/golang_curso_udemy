package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/routers"
	"webapp/src/utils"
)

//Exemplo de geracao de chave secret com securecookie
/*
func init() {
	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(string(hashKey))
	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(string(blockKey))

}
*/

func main() {

	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
	r := routers.Gerar()

	fmt.Printf("Executando webapp na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}

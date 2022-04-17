package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

/* Para gerar uma chave de 64 bits em base64
func init() {
	chave := make([]byte, 64)
	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	fmt.Println(stringBase64)

}
*/

func main() {

	config.Carregar()
	r := router.Gerar()
	fmt.Println("Escutando na porta " + fmt.Sprintf(":%d", config.Porta))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}

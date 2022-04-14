package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do módulo main!")
	auxiliar.Escrever()
	//auxiliar.escrever2() //funcoes só são exportadas para outros pacotes se iniciarem em maiúscula
	fmt.Println(checkmail.ValidateFormat("abc@"))
}

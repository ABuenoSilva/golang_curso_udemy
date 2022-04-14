package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 10
	if numero > 15 {
		fmt.Println("Maior do que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("outroNumero é maior que zero")
	} else if numero < -10 {
		fmt.Println("outroNumero é menor que -10")
	} else {
		fmt.Println("outroNumero é maior ou igual que -10 e menor ou igual a zero")
	}

	//Isso dá erro, if-init tem escopo
	//fmt.Println(outroNumero)

}

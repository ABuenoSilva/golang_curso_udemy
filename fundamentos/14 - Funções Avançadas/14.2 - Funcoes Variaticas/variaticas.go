package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

//Somente um parâmetro variático por função e tem que ser o último
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println(soma(1, 2, 3))
	fmt.Println(soma(1, 2, 5, 7))

	escrever("OIE!!", 1, 3, 4, 5)
}

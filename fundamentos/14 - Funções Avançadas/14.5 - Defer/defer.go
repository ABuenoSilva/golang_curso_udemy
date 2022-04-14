package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {

	defer fmt.Println("Media calculada. Retornando resultado")

	fmt.Println("Verificando aprovação")
	media := (n1 + n2) / 2

	if media >= 6 {
		// fmt.Println("Media calculada. Retornando resultado")
		return true
	}
	// fmt.Println("Media calculada. Retornando resultado")
	return false
}

func main() {
	/*
		defer funcao1() //DEFER - Adiar
		funcao2()
	*/
	fmt.Println(alunoEstaAprovado(7, 8))
}

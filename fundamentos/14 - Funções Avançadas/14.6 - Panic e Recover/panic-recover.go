package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Recovery executado com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6")
}

func main() {
	defer recuperarExecucao()
	fmt.Println(alunoEstaAprovado(6, 7))
	fmt.Println(alunoEstaAprovado(3, 7))
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução")

}

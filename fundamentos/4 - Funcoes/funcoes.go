package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//Posso utilizar um só tipo se todos forem iguais no parâmetro
//2 retornos em uma só função
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	fmt.Println(somar(10, 20))

	var f = func(txt string) string {
		fmt.Println("Função f")
		return "Função f: " + txt
	}
	resultado := f("meu texto")
	fmt.Println(resultado)

	//Resposta de funções precisam da mesmq qtde de variaveis
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// _ (underline) é usado pra ignorar um retorno
	resultadoCalculo1, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoCalculo1)

	_, resultadoCalculo2 := calculosMatematicos(10, 15)
	fmt.Println(resultadoCalculo2)

}

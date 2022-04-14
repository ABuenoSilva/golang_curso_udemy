package main

import "fmt"

//como o retorno tem nomes além dos tipos, não é necessário declarar as variáveis internas de retorno, e pode-se usar somente return no final
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	fmt.Println(calculosMatematicos(3, 5))
}

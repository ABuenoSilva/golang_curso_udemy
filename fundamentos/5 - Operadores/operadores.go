package main

import "fmt"

func main() {

	// ARITMETICOS

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 4
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 30

	//Isto dá erro, não pode fazer operações diretamente com tipos diferentes mesmo sendo ints
	//somaErrada := numero1 + numero2
	//Ou cria os dois com o mesmo tipo ou converte
	somaCerta := numero1 + int16(numero2)
	fmt.Println(somaCerta)

	//ATRIBUIÇÃO
	var variavel1 string = "string 1"
	variavel2 := "string 2"
	fmt.Println(variavel1, variavel2)

	//RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//LOGICOS
	verdadeiro, falso := true, false
	fmt.Println(".......")
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	//UNARIOS ( --var ou ++var não funcionam)
	numero := 10
	numero++
	fmt.Println(numero)
	numero += 15
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero -= 15
	fmt.Println(numero)
	numero *= 3
	fmt.Println(numero)
	numero /= 5
	fmt.Println(numero)
	numero %= 2
	fmt.Println(numero)

	//TERNARIO
	//Não funciona
	//texto := numero > 5 ? "Maior que 5" : "Menor qu e5"
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
}

package main

import (
	"errors"
	"fmt"
)

func main() {

	//Inteiros - se passar pelo máximo do tipo já dá erro em tempo de porgramacao
	var (
		_int8  int8  = 100
		_int16 int16 = 10000
		_int32 int32 = 1000000000
		_int64 int64 = 1000000000000000000
	)
	_int := 1000000000000000000 //Declaracao por inferencia sempre vai pegar "int" com o maior valor do hardware - no caso int64

	fmt.Println(_int8, _int16, _int32, _int64, _int)

	//uint - Inteiros sem sinal
	var numeroUning uint16 = 10
	fmt.Println(numeroUning)

	//alias
	//int32 = rune
	//int8 = byte
	var numero3 rune = 123456
	var numero4 byte = 255
	fmt.Println(numero3, numero4)

	//Reais - float32 e float64
	var numeroReal1 float32 = 123.45
	var numeroReal2 float64 = 5555464.55
	numeroReal3 := 5564.35 //Inferência pega o maior valor (32 ou 64) de acordo com o hardware

	fmt.Println(numeroReal1, numeroReal2, numeroReal3)

	//String - sempre aspas duplas, com aspas simples ele entende como char (numero)
	var str string = "Texto"
	str2 := "Texto2"

	fmt.Println(str, str2)

	char := 'a'
	fmt.Println(char)

	//Valor zero - sempre o valor inicial do tipo de dado, qdo declarado sem inicializar
	var texto string
	var num int16

	fmt.Println(texto, num)

	//Boolean (valor zero = false)
	var booleano1 bool
	booleano2 := true

	fmt.Println(booleano1, booleano2)

	//Erro (em go é um tipo - valor zer = <nil>)
	var erro error
	var erro2 error = errors.New("Novo erro")
	erro3 := errors.New("Outro erro")

	fmt.Println(erro, erro2, erro3)
}

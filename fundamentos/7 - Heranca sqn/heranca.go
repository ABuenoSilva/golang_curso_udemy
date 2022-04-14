package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	var p1 = pessoa{"Alexandre", "Bueno", 48, 182}
	fmt.Println(p1)
	var est1 = estudante{pessoa{"Zanza", "Garcia", 47, 160}, "Engenharia", "USP"}
	fmt.Println(est1)
	var est2 = estudante{p1, "Análise", "Fundação"}
	fmt.Println(est2)
	fmt.Println(est1.nome, est2.altura)

}

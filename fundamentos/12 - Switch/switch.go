package main

import "fmt"

func diaDaSemana(numero int) string {
	var diaDaSemana string
	switch numero {
	case 1:
		diaDaSemana = "Domingo"
		fallthrough
	case 2:
		diaDaSemana = "Segunda-Feira"
	case 3:
		diaDaSemana = "Terça-Feira"
	case 4:
		diaDaSemana = "Quarta-Feira"
	case 5:
		diaDaSemana = "Quinta-Feira"
	case 6:
		diaDaSemana = "Sexta-Feira"
	case 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Dia da semana inválido"
	}
	return diaDaSemana
}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-Feira"
	case numero == 3:
		return "Terça-Feira"
	case numero == 4:
		return "Quarta-Feira"
	case numero == 5:
		return "Quinta-Feira"
	case numero == 6:
		return "Sexta-Feira"
	case numero == 7:
		return "Sábado"
	default:
		return "Dia da semana inválido"
	}
}

func main() {
	dia := diaDaSemana(3)
	fmt.Println(dia)
	dia2 := diaDaSemana(1)
	fmt.Println(dia2)
}

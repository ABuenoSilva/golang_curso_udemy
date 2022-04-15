package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")

	/* laços anteriores
	i := 0
	for i < 10 {
		i++
		time.Sleep(time.Second / 2)
		fmt.Println(i)
	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		time.Sleep(time.Second / 2)
		fmt.Println(j)
	}
	*/
	nomes := [3]string{"Eu", "Você", "Eles"}

	for indice, nome := range nomes {
		fmt.Println(indice, ":", nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Felipe",
		"sobrenome": "Bueno",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//Loop infinito
	/*
		for {
			fmt.Println("Infinito")
		}
	*/
}

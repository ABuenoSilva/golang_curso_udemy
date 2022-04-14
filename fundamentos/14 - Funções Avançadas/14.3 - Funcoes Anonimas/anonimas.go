package main

import "fmt"

func main() {
	meuTexto := func(texto string) string {
		return fmt.Sprintf("Recebido => %s", texto)
	}("Parametro")

	fmt.Println(meuTexto)
}

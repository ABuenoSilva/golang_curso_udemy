package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Dentro do metodo salvar do usuario %s\n", u.nome)
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Alexandre", 48}
	usuario1.salvar()
	fmt.Println(usuario1)
	usuario1.fazerAniversario()
	fmt.Println(usuario1)
}

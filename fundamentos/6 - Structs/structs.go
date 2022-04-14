package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int16
}

func main() {
	fmt.Println("Arquivo structs")

	var enderecoEx endereco
	enderecoEx.logradouro = "Rua qualquer"
	enderecoEx.numero = 20

	var u usuario
	u.idade = 49
	u.nome = "Alexandre"
	u.endereco = enderecoEx
	fmt.Println(u)

	u2 := usuario{"Zanza", 48, endereco{"Avenida", 300}}
	fmt.Println(u2)

	u3 := usuario{idade: 25}
	fmt.Println(u3)

	u4 := usuario{"Zanza", 48, endereco{numero: 300}}
	fmt.Println(u4)

}

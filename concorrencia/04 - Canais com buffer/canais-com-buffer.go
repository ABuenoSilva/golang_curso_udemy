package main

import "fmt"

func main() {
	canal := make(chan string, 2) //Buffer de 2 canais, não vai bloquear a execução e dar deadlock até atingir a capacidade

	canal <- "Olá mundo"
	canal <- "Canal 2"
	//canal <- "deadlock!!!"

	mensagem := <-canal
	mensagem2 := <-canal
	//mensagem3 := <-canal //Outro tipo de deadlock
	fmt.Println(mensagem, mensagem2)
}

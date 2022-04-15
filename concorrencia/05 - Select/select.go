package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	//Essa abordagem faz o canal 1 perder tempo, pois tem que esperar o canal 2, que gera a cada 2s

	/* for errado
	for {
		mensagemCanal1 := <-canal1
		fmt.Println(mensagemCanal1)
		mensagemCanal2 := <-canal2
		fmt.Println(mensagemCanal2)

	}
	*/

	//Abordagem correta com select
	for {
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}
}

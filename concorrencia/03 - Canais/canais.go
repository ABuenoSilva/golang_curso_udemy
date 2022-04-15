package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("OLÁ MUNDÃO!!!", canal)

	//Tomar cuidado para verificar se o canal está aberto, senão pode dar deadlock e só vai aparecer em tempo de execução
	/*
		for {
			mensagem, aberto := <-canal
			if !aberto {
				break
			}
			fmt.Println(mensagem)
		}
	*/
	//Versão refatorada do for
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}

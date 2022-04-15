package main

import (
	"fmt"
	"time"
)

func main() {
	//CONCORRÊNCIA != PARALELISMO
	go escrever("OLÁ MUNDÃO !!!") //goroutine
	escrever("PROGRAMANDO EM GO!!!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

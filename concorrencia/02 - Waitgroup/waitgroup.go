package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("OLÁ MUNDÃO !!!")
		waitGroup.Done()
	}()

	go func() {
		escrever("PROGRAMANDO EM GO!!!")
		waitGroup.Done()
	}()

	waitGroup.Wait()
	fmt.Println("Finalizado!!")
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

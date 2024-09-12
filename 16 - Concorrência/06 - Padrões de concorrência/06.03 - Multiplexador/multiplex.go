package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplex(escrever("All the single ladies"), escrever("Now put your hands up"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(text string) <-chan string {
	canal := make(chan string)
	iteration := 1

	go func() {
		for {
			canal <- fmt.Sprintf("%d - Received value as: %s", iteration, text)
			iteration++
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}

func multiplex(canalDeEntrada_01, canalDeEntrada_02 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func ()  {
		for {
			select{
			case message := <-canalDeEntrada_01:
				canalDeSaida <- message
			case message := <-canalDeEntrada_02:
				canalDeSaida <- message
			}
		}
	}()

	return canalDeSaida
}

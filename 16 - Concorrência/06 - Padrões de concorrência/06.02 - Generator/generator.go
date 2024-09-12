package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("All the single ladies")

	for i := 0; i < 12; i++ {
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
			time.Sleep(time.Second * 5)
		}
	}()

	return canal
}

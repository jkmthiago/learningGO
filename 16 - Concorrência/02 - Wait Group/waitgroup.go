package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func(){
		escrever("Texto 01 Número")
		waitGroup.Done() // -1
	}()

	go func(){
		escrever("Texto 02 Número")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // goroutines := 0
}

func escrever(texto string) {
	for i:=0; i < 5; i++{
		fmt.Println(texto, i+1)
		time.Sleep(time.Second)
	}	
}
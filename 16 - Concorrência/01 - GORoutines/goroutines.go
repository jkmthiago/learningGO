package main

import (
	"fmt"
	// "time"
)

func main() {
	go escrever("Texto 01 Número")
	escrever("Texto 02 Número")
}

func escrever(texto string) {
	for i:=0; true; i++{
		fmt.Println(texto, i)
		// time.Sleep(time.Second)
	}	
}
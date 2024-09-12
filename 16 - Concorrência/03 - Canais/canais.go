package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Texto Número ", canal)

	// Uma forma de fazer o recebimento da informação do canal
	// for {
	// 	answer, isChannelOpen := <- canal

	// 	if !isChannelOpen {
	// 		break
	// 	}

	// 	fmt.Println(answer)
	// }

	// Forma mais direta e correta
	for answer := range canal {
		fmt.Println(answer)
	}
}

func escrever(texto string, canal chan string) {
	for i:=0; i < 5; i++{
		canal <- texto + strconv.Itoa(i+1)
		time.Sleep(time.Second)
	}	
	
	close(canal)
}
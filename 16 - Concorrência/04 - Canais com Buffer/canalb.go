package main

import "fmt"

//canal com buffer somente é bloqueante quando atingido a capacidade máxima do canal.

func main() {
	canal := make(chan string, 2) // canal com buffer de envio de até 2 mensagens
	canal <- "Olá canal!"
	canal <- "Show de bola?"

	mensagem_01 := <- canal
	mensagem_02 := <- canal

	fmt.Println(mensagem_01,mensagem_02)
}
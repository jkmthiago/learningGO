package main

import (
	"fmt"
	"time"
)

func main() {
	canal_01, canal_02 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal_01 <- "Canal 01"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal_02 <- "Canal 02"
		}
	}()

	for {
		select {
		case mensagem_canal_01 := <-canal_01:
			fmt.Println(time.Now(), "-", mensagem_canal_01)
		case mensagem_canal_02 := <-canal_02:
			fmt.Println(time.Now(), "-", mensagem_canal_02)
		}
	}
}

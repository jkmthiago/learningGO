package main

import (
	"fmt"
	"log"
	"os"
	"shelApplication/app"
)

func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatalln(erro)
	}
}
package main

import (
	"fmt"
	// "time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i + 1)
		// time.Sleep(time.Second)
	}

	numeros := [3]int8{1,2,3}

	for indice, numero := range numeros {
		fmt.Println(indice + 1, "-", numero)
	}

	usuario := map[string]string{
		"Nome" : "Thiago",
		"Sobrenome" : "Eleuterio",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, "-", valor)
	}

	i := 1
	var stop bool	
	for !stop {	
		if i != 10 {
			fmt.Println(i)
			i++		
		} else { 
			fmt.Println("Limite atingido")
			stop = true
		}
	}
}
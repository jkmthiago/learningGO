package main

import "fmt"

func closure() func() {
	var texto string = "Dentro da Função Closure"

	funcao := func ()  {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da Função Main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
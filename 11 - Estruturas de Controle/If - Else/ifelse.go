package main

import "fmt"

func main() {
	numero :=10.8

	if numero < 10.8 {
		fmt.Println("O número é menor que 10.8")
	} else {
		fmt.Println("O número é maior ou igual a 10.8")
	}

	if sim := numero; sim < 10.8 {
		fmt.Println("Menor que 10.8")
	} else if sim > 10.8 {
		fmt.Println("Maior que 10.8")
	} else {
		fmt.Println("Igual a 10.8")
	}
}
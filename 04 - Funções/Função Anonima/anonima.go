package main

import "fmt"

func main() {
	
	retorno := func(numero int8)(theOne string){
		fmt.Println("Olá mundo")
		if numero == 1 {
			theOne = "The One Above All"
			return fmt.Sprintf("Retornando o valor: %s", theOne)
		} else {
			return fmt.Sprintf("Entrada Inválida")
		}
	}(1)

	fmt.Println(retorno)
}
package main

import "fmt"

func invertTheSignal(numero int)int {
	return numero * -1
}

func invertSignalWithPointers(numero *int) {
	*numero *= -1
}

func main() {
	numero := 20
	numeroInvertido := invertTheSignal(numero)
	
	fmt.Println("_____________________")
	fmt.Println("Numero          : ", numero)
	fmt.Println("Numero invertido:", numeroInvertido)
	fmt.Println("_____________________")
	
	novoNumero := 50
	fmt.Println("_____________________")
	fmt.Println("Numero          : ", novoNumero)
	invertSignalWithPointers(&novoNumero)
	fmt.Println("Numero invertido:", novoNumero)
	fmt.Println("_____________________")
}
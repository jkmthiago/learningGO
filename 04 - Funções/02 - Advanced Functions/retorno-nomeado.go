package main

import "fmt"

func calculadora(n1, n2 int8) (soma int8, sub int8, mult int8, div float32){
	soma = n1 + n2
	sub = n1 - n2
	mult = n1 * n2
	div = float32(n1) / float32(n2)  
	return
}

func main() {
	soma, sub, mult, div := calculadora(40, 21)
	fmt.Println("O Resultado das operações aritiméticas usando a função nomeada é?:")
	fmt.Println("Soma:", soma)
	fmt.Println("Sub:", sub)
	fmt.Println("Mult:", mult)
	fmt.Println("Div:", div)
}
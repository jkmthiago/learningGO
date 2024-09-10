/* Esse código é uma representação da declaração e uso de funções
   na linguagem go. */

package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculadora (n1 int8, n2 int8)(int8, int8, int8, float32){
	soma := n1 + n2
	sub := n1 - n2
	mult := n1 * n2
	div := float32(n1)/float32(n2)

	return soma, sub, mult, div
}

func main() {
	// soma := somar(23,25)
	// fmt.Println("O resultado da soma é", soma)

	var f = func (txt string) string  {
		fmt.Println(txt)
		return txt		
	}

	fmt.Println("Olha a função", f("f"), "está funcionndo")

	var n1, n2 int8 = 24, 35

	fmt.Println("Vamos somar, subtrair, multiplicar e dividir dois numeros sendo eles", n1, "e", n2)

	soma, sub, mult, div := calculadora(n1, n2)
	
	fmt.Println("Aqui abaixo segue o resultado")
	fmt.Println("Soma:", soma)
	fmt.Println("Subtração:", sub)
	fmt.Println("Multiplicação:", mult)
	fmt.Println("Divisão:", div)
}
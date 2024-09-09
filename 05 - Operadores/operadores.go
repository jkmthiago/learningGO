/* Esse código é uma representação do uso de operadores na
   linguagem go. */

package main

import "fmt"

// OPERADORES ARITMÉTICOS - LEMBRAR DE NÃO REALIZAR OPERAÇÃO DE TIPOS DIFERENTES
func calculadora (n1, n2 int8)(int8, int8, int8, float32, int8){
	soma := n1 + n2
	sub := n1 - n2
	mult := n1 * n2
	div := float32(n1)/float32(n2)
	resto := n1%n2
	return soma, sub, mult, div, resto
}

// Operadores Relacionais
func comparaAE(n1, n2 int8) {
	fmt.Println(n1, "é maior que", n2, "         :", n1 >  n2)
	fmt.Println(n1, "é maior ou igual que", n2, ":", n1 >= n2)
	fmt.Println(n1, "é menor que", n2, "         :", n1 <  n2)
	fmt.Println(n1, "é menor ou igual que", n2, ":", n1 <= n2)
	fmt.Println(n1, "é igual que", n2, "         :", n1 == n2)
	fmt.Println(n1, "é diferente que", n2, "     :", n1 != n2)
}

func main() {
	var n1, n2 int8 = 24, 35

	fmt.Println("Vamos somar, subtrair, multiplicar e dividir dois numeros sendo eles", n1, "e", n2)

	soma, sub, mult, div, resto := calculadora(n1, n2)
	
	fmt.Println("Aqui abaixo segue o resultado")
	fmt.Println("Soma:", soma)
	fmt.Println("Subtração:", sub)
	fmt.Println("Multiplicação:", mult)
	fmt.Println("Divisão:", div)
	fmt.Println("Resto:", resto)

	fmt.Println("Agora vamos compará-los")

	comparaAE(n1, n2)
}
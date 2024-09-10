package main

import "fmt"

func soma(numeros ...int){
	var total int
	for _, numero := range numeros {
		fmt.Println("Total parcial:", total)
		total += numero
	}
	fmt.Println("Total final:", total)
}

func main() {
	soma(2,3,4,5,7,8,9)
}
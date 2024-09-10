package main

import "fmt"

type pessoa struct {
	primeiro_nome string
	sobrenome string
	idade uint8
	altura	uint8
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main() {
	p1 := pessoa{"Thiago", "Eleuterio da Silva", 23, 170}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia da Computação", "Faculdade de Computação"}
	fmt.Println(e1)

	e2 := estudante{pessoa{"Henrique", "Pereira Viana", 23, 175}, "Engenharia da Computação", "Faculdade de Computação"}
	fmt.Println(e2)
}
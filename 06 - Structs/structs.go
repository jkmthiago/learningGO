package main

import "fmt"

type user struct {
	nome  string
	idade uint8
	cargo string
}

func main() {
	var u user
	u.nome = "Thiago Eleuterio da Silva"
	u.idade = 23
	u.cargo = "Assistente de sistemas"

	fmt.Println(u)

	u2 := user{"Henrique Pereira Viana", 23, "Assistente de Suporte"}
	 
	fmt.Println(u2)

	u3 := user{idade: 25}
	fmt.Println(u3)
}


	
package main

import "fmt"

func main() {
	usuario := map[string]string {
		"Nome": "Thiago", 
		"Sobrenome" : "Eleuterio", 
	}

	fmt.Println(usuario["nome"])

	registro := map[string]map[string]string {
		"usuario_registro": usuario,
		"cadastro_registro": {
			"e-mail": "thiago-eleuterio@teste123.com",
		},
	}

	fmt.Println(registro)
}
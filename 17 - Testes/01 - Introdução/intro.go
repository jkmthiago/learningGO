package main

import (
	"fmt"
	"testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rua dos Bobos")
	fmt.Println(tipoEndereco)
}
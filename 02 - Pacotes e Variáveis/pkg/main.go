/* Esse código é uma representação da declaração de pacotes na
   linguagem go e como utilizá-los de forma local ou remota. */

package main

import (
	"fmt"
	"modules/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("The one above all is only someone as big as you who sees you bellow him")	
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("thiago@hotmail.com")
	fmt.Println(erro)
}
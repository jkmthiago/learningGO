/* Esse código é uma representação dos tipos de dados na linguagem go. */

package main

import (
	"fmt"
)

func main() {
	var oito_01 int8 = -83
	var oito_02 byte = 'a'
	var char byte = 'd'
	var dezesseis_01 int16 = -6464
	var trintaEDois_01 int32 = -827498324
	var trintaEDois_02 rune = -32423423
	var trintaEDois_03 rune = 's'
	var sessentaEQuatro_01 int64 = -932840238482389432
	var oito_03 uint8 = 83
	var dezesseis_02 uint16 = 62264
	var trintaEDois_04 uint32 = 927498324
	var sessentaEQuatro_02 uint64 = 2409238439877979879
	var real_32 float32 = 8.2345678
	var real_64 float64 = 4454654654564.46545645646464654564654
	var nome string = "thiago"

	fmt.Println(oito_01, oito_02, oito_03, char, dezesseis_01, dezesseis_02, trintaEDois_01, trintaEDois_02, trintaEDois_03, trintaEDois_04, sessentaEQuatro_01, sessentaEQuatro_02, real_32, real_64, nome)
}
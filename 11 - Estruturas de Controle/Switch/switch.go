package main

import "fmt"

func diaDaSemana1(numero int) string{
	switch numero {
	case 0:
		return "Domingo"
	case 1:
		return "Segunda-feira"
	case 2:
		return "Terça-feira"
	case 3:
		return "Quarta-feira"
	case 4:
		return "Quinta-feira"
	case 5:
		return "Sexta-feira"
	case 6:
		return "Sábado"
	default:
		return "Número Invalido"
	}
}

func diaDaSemana2(numero int) string{
	switch  {
	case numero == 0:
		return "Domingo"
	case numero == 1:
		return "Segunda-feira"
	case numero == 2:
		return "Terça-feira"
	case numero == 3:
		return "Quarta-feira"
	case numero == 4:
		return "Quinta-feira"
	case numero == 5:
		return "Sexta-feira"
	case numero == 6:
		return "Sábado"
	default:
		return "Número Invalido"
	}
}

func diaDaSemana3(numero int) string{
	var diaDaSemana string
	switch  {
	case numero == 0:
		diaDaSemana = "Domingo"
	case numero == 1:
		diaDaSemana = "Segunda-feira"
	case numero == 2:
		diaDaSemana = "Terça-feira"
	case numero == 3:
		diaDaSemana = "Quarta-feira"
	case numero == 4:
		diaDaSemana = "Quinta-feira"
	case numero == 5:
		diaDaSemana = "Sexta-feira"
	case numero == 6:
		diaDaSemana = "Sábado"
		fallthrough
	default:
		diaDaSemana = "Número Invalido"
	}

	return diaDaSemana
}

func main() {
	dia1 := diaDaSemana1(6)
	dia2 := diaDaSemana2(3)
	dia3 := diaDaSemana3(6)

	fmt.Println(dia1, dia2, dia3)
}
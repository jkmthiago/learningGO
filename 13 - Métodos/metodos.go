package main

import (
    "fmt" 	
    "time"
)

type usuario struct {
	nome string
	idade uint8
	nasc_dia int
	nasc_mes int
	nasc_ano int
}

func (u usuario) salvar() {
	fmt.Printf("\nSaving data from User %s on dadabase\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u usuario) imprimirUsuario() {
	fmt.Println("__________________________")
	fmt.Printf("Nome:           %s\nIdade:          %d\n", u.nome, u.idade)	
	fmt.Println("Maior de idade:", u.maiorDeIdade())
	fmt.Printf("Nascimento    :%d/%d/%d\n", u.nasc_dia, u.nasc_mes, u.nasc_ano)
	fmt.Println("__________________________")
}

func whichMonth(month string) int {
	switch month {
		case "January":	
			return 1
		case "February":
			return 2
		case "March":
			return 3
		case "April":
			return 4
		case "May":
			return 5
		case "June":
			return 6
		case "July":
			return 7
		case "August":
			return 8
		case "September":
			return 9
		case "Octuber":
			return 10
		case "November":
			return 11
		case "December":
			return 12	
		default:
			return 0	
	}	
}

func (u *usuario) birthday() bool {
	now := time.Now()
	month := whichMonth(now.Month().String())
	if month == u.nasc_mes && now.Day() == u.nasc_dia {
		fmt.Println("Feliz aniversário!!!")
		fmt.Println("Hoje você completa", now.Year()-u.nasc_ano, "anos")
		u.idade++
		fmt.Println("Sua nova idade é:", u.idade)
		return true
	}
	return false
}

func main() {
	fmt.Println("Criando o usuário")
	user_teste := usuario{"Thiago", 20, 10, 9, 2001}
	user_teste.imprimirUsuario()
	user_teste.salvar()
	fmt.Println("É aniversário do Usuário: ", user_teste.birthday())
}
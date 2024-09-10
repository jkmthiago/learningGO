package main

import "fmt"

func recoverOperation () {
	if r := recover(); r != nil {
		fmt.Println("Tentando recuperar a execução")
	}
}

func studentPerformance(n1, n2 float32) (bool) {
	defer recoverOperation()
	defer println("O Desempenho do aluno foi")
	fmt.Println("Analisando performance do aluno")
	media := (n1 + n2) / 2

	fmt.Println("Media do aluno:", media)
	
	if media > 7 {
		return true		
	} else if media < 7 {
		return false
	}

	panic("Media igual a 7")
}

func main() {
	isTheStudentApproved := studentPerformance(1.8, 6.2)
	if isTheStudentApproved {
		fmt.Println("Estudante aprovado")
	} else {
		fmt.Println("Estudante reprovado")
	}
}
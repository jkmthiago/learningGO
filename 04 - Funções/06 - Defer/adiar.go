package main

import "fmt"

func studentPerformance(n1, n2 float32) (bool) {
	defer println("O Desempenho do aluno foi")
	fmt.Println("Analisando performance do aluno")
	media := (n1 + n2) / 2

	fmt.Println("Media do aluno:", media)
	
	if media >= 7 {
		return true		
	} else {
		return false
	}
}

func main() {
	isTheStudentApproved := studentPerformance(7.8, 5.2)
	if isTheStudentApproved {
		fmt.Println("Estudante aprovado")
	} else {
		fmt.Println("Estudante reprovado")
	}
}
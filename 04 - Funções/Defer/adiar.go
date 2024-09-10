package main

import "fmt"

func studentPerformance(n1, n2 float32) (approved bool) {
	fmt.Println("Analisando performance do aluno")
	media := (n1 + n2) / 2

	fmt.Println("Media do aluno:", media)
	
	if media >= 7 {
		approved = true		
	} else {
		approved = false
	}
	return	
}

func main() {
	isTheStudentApproved := studentPerformance(7.8, 6.2)
	if isTheStudentApproved {
		fmt.Println("Estudante aprovado")
	} else {
		fmt.Println("Estudante reprovado")
	}
}
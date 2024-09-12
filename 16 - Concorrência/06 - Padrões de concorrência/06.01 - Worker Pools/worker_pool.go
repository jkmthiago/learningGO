package main

import "fmt"

func fibonacci(posicao uint) uint {

	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao - 2) + fibonacci(posicao - 1)
}

func worker(toBeCalculated <-chan uint, calculatedResults chan<- uint) {
	for position := range toBeCalculated {
		calculatedResults <- fibonacci(position)
	}
}

func main() {
	toBeCalculated := make(chan uint, 45)
	calculatedResults := make(chan uint, 45)

	go worker(toBeCalculated, calculatedResults)
	go worker(toBeCalculated, calculatedResults)

	for i := 0; i < 45; i++ {
		toBeCalculated <- uint(i)	
	}
	close(toBeCalculated)

	for i := 0; i < 45; i++ {
		results := <- calculatedResults
		fmt.Println(i+1, "-", results)
	}
}
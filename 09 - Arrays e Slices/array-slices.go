package main

import "fmt"

func main() {
	// 01 - Arrays e Slice
	var array1[5] int
	array1[0] = 1

	fmt.Println(array1)

	array2 := [...]int {1,2,3,4,5,6}
	
	fmt.Println(array2)

	slice := []int {1,2,3,4,5,6}
	
	fmt.Println(slice)
	
	slice = append(slice, 7, 8, 9, 10)
	
	fmt.Println(slice)

	// 02 - Arrays internos
	slice3 := make([]float32, 10, 15)

	fmt.Println(slice3)
}
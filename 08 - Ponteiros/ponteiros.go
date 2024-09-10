package main

import "fmt"

func main() {

	var var1 int8 = 10
	var var2 int8 = var1

	fmt.Println(var1, var2)

	var1 = 12

	fmt.Println(var1, var2)
	
	var var3 int8 = 10
	var var4 *int8 = &var3

	fmt.Println(var3, *var4)
	
	var3 = 12

	fmt.Println(var3, *var4)
}
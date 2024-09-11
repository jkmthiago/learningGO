package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("Thiago")
	generica(2)
	generica(3.25)
	generica(true)
}
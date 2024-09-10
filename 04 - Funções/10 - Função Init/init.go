package main

import "fmt"

var thef string = "The F***"

func init() {
	thef += " Is Wrong With You"
}

func main() {
	fmt.Println(thef)
}
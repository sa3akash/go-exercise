package main

import "fmt"

func main() {
	var a int = 10
	var b float64 = 20.5
	var c string = "Go Programming"

	// shorthand declaration
	d := 30
	name := "Gopher"
	age := 15

	fmt.Println("Shorthand Integer:", d)
	fmt.Println("Shorthand String:", name)
	fmt.Println("Shorthand Age:", age)

	fmt.Println("Integer:", a)
	fmt.Println("Float:", b)
	fmt.Println("String:", c)
}
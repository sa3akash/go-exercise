package main

import "fmt"

func main() {
	const name string = "Gopher"
	const pi float64 = 3.14159
	const isGoFun bool = true

	// user can not change the value of a constant
	// name = "Gopher"

	fmt.Println("Name:", name)

	// constants group

	const (
		country    string  = "GoLand"
		version    float64 = 1.18
		isAwesome  bool    = true
	)

	fmt.Println("Country:", country)
	fmt.Println("Version:", version)
	fmt.Println("Is Go Awesome?:", isAwesome)


}
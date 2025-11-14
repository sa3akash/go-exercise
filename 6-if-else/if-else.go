package main

import "fmt"

func main() {
	num := 2

	if num > 5 {
		fmt.Println(num, "is greater than 5")
	} else {
		fmt.Println(num, "is not greater than 5")
	}

	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}
}
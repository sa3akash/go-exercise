package main

import "fmt"

func main() {

	// traditional for loop
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// while loop using for
	i := 1
	for i <= 5 {
		fmt.Println("While Loop Iteration:", i)
		i++
	}

	// infinite loop with break
	j := 1
	for {
		if j > 5 {
			break
		}

		fmt.Println("Infinite Loop Iteration:", j)
		j++
	}

	// continue statement
	for k := 1; k <= 5; k++ {
		if k%2 == 0 {
			continue
		}

		fmt.Println("Odd Number:", k)
	}

	// range loop
	for i := range []int{10, 20, 30, 40, 50} {
		fmt.Println("Range Loop Index:", i)
	}

}

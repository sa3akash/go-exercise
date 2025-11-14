package main

import "time"

func main() {
	i := 5

	switch i {
	case 1:
		println("One")
	case 2:
		println("Two")
	case 3:
		println("Three")
	case 4:
		println("Four")
	case 5:
		println("Five")
	default:
		println("Unknown number")
	}

	// multiple condition switch
	switch i {
	case 1, 3, 5, 7, 9:
		println("Odd number")
	case 2, 4, 6, 8, 10:
		println("Even number")
	default:
		println("Number out of range")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		println("It's the weekend!")
	default:
		println("It's a weekday.")
	}

	// type switch
	var x interface{} = 42

	switch v := x.(type) {
	case int:
		println("x is an integer:", v)
	case string:
		println("x is a string:", v)
	default:
		println("Unknown type")
	}

	whoAmi := func(i any) {
		switch v := i.(type) {
		case int:
			println("I am an integer:", v)
		case string:
			println("I am a string:", v)
		default:
			println("I am of unknown type")
		}
	}

	whoAmi(100)
	whoAmi("Hello, Go!")
	whoAmi(3.14)
}

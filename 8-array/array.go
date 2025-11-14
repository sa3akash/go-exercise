package main

import "fmt"

func main() {
	// This is a placeholder for the main function.

	var nums [5]int
	nums[0] = 10
	nums[1] = 20
	nums[2] = 30
	nums[3] = 40
	nums[4] = 50

	fmt.Println(len(nums))
	fmt.Println(nums)

	// slice example
	slice := []string{"apple", "banana", "cherry", "date", "elderberry"}
	fmt.Println(len(slice))
	fmt.Println(slice)

}
package main

import "fmt"

func main() {

	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("Original slice:", slice)
	fmt.Println("Length of slice:", len(slice))

	// Slicing the slice
	subSlice1 := slice[2:5] // elements from index 2 to 4
	fmt.Println("Sub-slice from index 2 to 4:", subSlice1)

}

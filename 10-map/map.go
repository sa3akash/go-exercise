package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	for key, value := range m {
		fmt.Println(key, value)
	}

	delete(m, "two")

	fmt.Println("Map length:", len(m))

	newMap := map[string]int{
		"alpha": 1,
		"beta":  2,
		"gamma": 3,
		"delta": 4,
	}

	newMap["four"] = 4

	fmt.Println("New Map:", newMap)

	for key, value := range newMap {
		if key == "beta" {
			fmt.Println(key, value)
		}
	}

}
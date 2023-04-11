package main

import "fmt"

func main() {
	// Passing slice as argument
	//  Can give unlimited arguments, as its a slice it will expand
	proRes := proAdder(2, 5, 7, 3, 5, 9)
	fmt.Println("PRO result is",proRes)
}
func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

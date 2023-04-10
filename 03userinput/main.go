package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome ")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")
	//  Comma ok syntax or error ok syntax
	//  Its like try and catch in javascript, the first part being the try and the second part, catch
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating, ",input)
}
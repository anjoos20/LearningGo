package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome ")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")
	//  Comma ok syntax or error ok syntax
	//  Its like try and catch in javascript, the first part being the try and the second part, catch
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating, ",input)
	var numRating, err  = strconv.ParseFloat(strings.TrimSpace(input),64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating + 1)
	}
}
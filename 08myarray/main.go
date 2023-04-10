package main

import "fmt"

func main() {
	//Ifusingarray,needtoprovidethesize
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Orange"
	// Here when we print the data at index=2 is space
	fmt.Println("FruitList is ",fruitList)
	fmt.Println("Length of FruitList is ",len(fruitList))

	var vegList = [5]string{"beans","peas","potatoes"}
	//  The declared value(i.e. 5 ) is the length here
	fmt.Println("Vegy list length is: ", len(vegList))

}
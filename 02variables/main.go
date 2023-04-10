package main

import "fmt"

func main() {
	var userName string = "Anju Balakrishnan"
	fmt.Println(userName)
	fmt.Printf("Variable is of type : %T \n", userName)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

//  If value is 256, it gives an error of numeric overflow	
	var smallVal uint8  = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)

	// for float32 value printed is 255.44432
	// for float64 value printed is 255.44432789453316
	// i.e. more precision
	var smallFloat float64  = 255.444327894533166345
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)
}
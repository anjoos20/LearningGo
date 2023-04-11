package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This needs to go in a file"
	file, err := os.Create("./myfile.txt")
	if err != nil {
		// Panic will shut down the execution of the program and will show you what you are facing
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		// Panic will shut down the execution of the program and will show you what you are facing
		panic(err)
	}
	fmt.Println("Length is: " , length)
	readFile("./myfile.txt")
	defer file.Close()
}

func readFile(fileName string){
	// Normally, when we read the file, it will be in the byte format
	databyte,err := ioutil.ReadFile(fileName)
	// error handling function
	checkNilErr(err)
	fmt.Println("Text data iside the file is: \n " , string(databyte))
}

func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}
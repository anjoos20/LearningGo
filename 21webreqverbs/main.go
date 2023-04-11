package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	// PerformGetRequest()
	PerformPostJsonRequest()
}
// Function which performs the get request
func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println()
	fmt.Println("StatusCode: ", response.StatusCode)
	// fmt.Println("Content length is: ", response.ContentLength)
	// Create a builder
	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount,_ := responseString.Write(content)
	// This is the same as content length
	fmt.Println("bytecount: ",byteCount)
	fmt.Println("response: ",responseString.String())
	// fmt.Println(string(content))
	// Content is in the byte format
}

func PerformPostJsonRequest(){
	const myurl = "http://localhost:8000/post"

	// Fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price": 0,
			"platform":"learnCodeOnline.in"
		}
	`)
	response, err := http.Post(myurl, "application/json",requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
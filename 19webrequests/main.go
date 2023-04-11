package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://openai.com/blog/chatgpt"

func main() {
	// The provided code is a simple Go program that makes an HTTP GET request to a specified URL
	//  and reads the response body.
	response, err :=  http.Get(url)
	checkNilErr(err)
	// Response is of type: *http.Response
	// i.e. we get the original reference of the response
	fmt.Printf("Response is of type: %T\n", response)
	// Its the callers responsibility to close the connection in Go
	defer response.Body.Close()
	databytes, err := io.ReadAll(response.Body)
	checkNilErr(err)
	content := string(databytes)
	fmt.Println("content is",content)
}

func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}
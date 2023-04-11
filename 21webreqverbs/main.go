package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	PerformGetRequest()
}
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
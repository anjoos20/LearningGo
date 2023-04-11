package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=hfghghvx"

func main() {
	fmt.Println(myurl)
	// Parsing the url
	result,_ := url.Parse(myurl)
	fmt.Println("scheme",result.Scheme)
	fmt.Println("Host",result.Host)
	fmt.Println("Path",result.Path)
	fmt.Println("RawQuery",result.RawQuery)
	// Port is not a propert but a method
	fmt.Println("Port",result.Port())
	qparams := result.Query()
	// type is url.Values means key-value pairs
	fmt.Printf("The type of query params are: %T", qparams)
	fmt.Println(qparams["coursename"])
	for key,val := range qparams {
		fmt.Println("Param is:", key,val)
	}
	//  building the url from the parts
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "tutcss",
		RawPath: "user=anju",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
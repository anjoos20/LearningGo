package main

import "fmt"

func main() {
	// There is no inheritance in go. Also no super / parent
	anju := User{"Anju","anju@gmail.com",true,25}
	fmt.Println(anju)
	// for struct if %+v, it includes key as well
	fmt.Printf("anju details are: %+v\n", anju)
	fmt.Printf("Name is %v and email is %v", anju.Name, anju.Email)
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}
package main

import "fmt"

func main() {
	// There is no inheritance in go. Also no super / parent
	anju := User{"Anju","anju@gmail.com",true,25}
	fmt.Println(anju)
	// for struct if %+v, it includes key as well
	fmt.Printf("anju details are: %+v\n", anju)
	fmt.Printf("Name is %v and email is %v\n", anju.Name, anju.Email)
	// Using the method
	anju.GetStatus()
	anju.NewMail()
	// when the struct is printed again it retains the old value i.e. in a function, copy of the argument is passed
	// And that is the usecase of pointers
	fmt.Printf("details after NewMail are: %+v\n", anju)
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
	// This particular property is not exportable
	// oneAge int
}
// Method
func (u User) GetStatus(){
	fmt.Println("Is User active:", u.Status )	
}
// Here a copy is passed as an argument
func (u User) NewMail(){
	u.Email = "test@gmail.com"
	fmt.Println("Email of this user is:",u.Email)
}
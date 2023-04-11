package main

import "fmt"

func main() {
	days := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	fmt.Println(days)

	//  type 1
	// But, in a case where you know, you want to list through the entire loop you wont be using
	// for d:=0; d< len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// type 2
	// for i := range days {
	// 	fmt.Println(days[i])
	// }
	// type 3
	// for index, day := range days{
	// 	fmt.Printf("Index is %v and value is %v\n", index, day)
	// }
	rogueValue := 1
	for rogueValue < 10 {

		if rogueValue == 2{
			goto test
		}
		if rogueValue == 5{
			// break
			rogueValue++
			continue
		}
		fmt.Println("value is", rogueValue)
		rogueValue++
	}

	test:
		fmt.Println("Jumping out of the loop")
}
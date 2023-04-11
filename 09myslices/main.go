package main

import (
	"fmt"
	"sort"
)

func main() {
	//  No need to specify size
	// We can add as many values as we need and it automatically expands the list
	var fruitList = []string{"apple", "orange", "peach"}
	fmt.Printf("The type of fruitList is %T", fruitList)
	// adding values
	// Append va;ues to fruitList and then move the final product to fruitList
	fruitList = append(fruitList,"Mango","Banana")
	// Means from position 1 to default i.e. all values
	//  fruitList = append(fruitList[1:])
	// fmt.Println(fruitList)
	// Means from position 1 to 2 i.e the range is noninclusive
	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)
	// Means from position 0 to 3 but 3 noninclusive
	fruitList = append(fruitList[:3])

	fmt.Println(fruitList)
	// To allocate memory new and make keywords
	// Lets utilise make to define a slice

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 345
	highScore[2] = 456
	highScore[3] = 567
	// Gonna give error that index out of boun
	// highScore[4] = 555
	// But with append method, it allocates more memory and can include more values
	highScore = append(highScore, 555, 666, 777)
	sort.Ints(highScore)

	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))

	//  To remove a value from slice based on the index

	var courses = []string{"reactjs","javascript","swift","python","ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)
}
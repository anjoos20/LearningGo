package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["JS"] = "javascript"
	languages["RB"] = "ruby"
	languages["PY"] = "Python"
	fmt.Println("list of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])
	//  Delete a single entry from the map
	delete(languages,"RB")
	fmt.Println("list of all languages after delete: ", languages)
	//  Looping thru maps
	for key, value := range languages {
		fmt.Println(key, "=>" , value)
	}
}
package main

import "fmt"

func main() {
	// init with key/value pairs
	dict := map[string]string{"red": "#da1337", "orange": "#e95a22"}
	fmt.Println(dict)

	// init with make()
	dict_m := make(map[string]int)
	fmt.Println(dict_m)
	dict_m["blue"] = 4444
	fmt.Println(dict_m)

	// nil map
	colors := map[string]string{}
	colors["red"] = "#da1337"
	fmt.Println(colors)

	var colors2 map[string]string
	//colors2["red"] = "#da1337" // nil map can't save key/value
	fmt.Println(colors2)

	// check existence
	value, exists := colors["red"]
	fmt.Println(value, exists)
	value, exists = colors["blue"]
	fmt.Println(value, exists)
}

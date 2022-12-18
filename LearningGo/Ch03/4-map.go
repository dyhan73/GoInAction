package main

import (
	"fmt"
)

func defineMap() {
	//var nilMap map[string]int{"key1":1}
	//nilMap["key2"] = 2 // panic
	//fmt.Println(nilMap)

	totalWins := map[string]int{}
	totalWins["team1"] = 4
	totalWins["team2"] = 5
	fmt.Println(totalWins)

	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams)

	ages := make(map[string]int, 5)
	ages["han"] = 49
	fmt.Println(ages)
}

func readWriteMap() {
	// write and read map
	totWins := map[string]int{}
	totWins["Orcas"] = 1
	totWins["Lions"] = 2
	fmt.Println(totWins["Orcas"])
	fmt.Println(totWins["Kittens"])
	totWins["Kittens"]++
	fmt.Println(totWins["Kittens"])
	totWins["Lions"] = 3
	fmt.Println(totWins["Lions"])
}

func checkExistenceWithCommaOk() {
	// check existence with comma OK
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)

	// delete map
	delete(m, "hello")
	fmt.Println(m)
}

func usingMapAsSet() {
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}
	if intSet[10] {
		fmt.Println("10 is in the set")
	}
}

//	func usingMapAsSetWithStruct() {
//		intSet := map[int]struct{}{}
//		vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
//		for _, v := range vals {
//			intSet[v] = struct{} // this line is wrong in v1.19
//		}
//		if _, ok := intSet[5]; ok {
//			fmt.Println("5 is in the set")
//		}
//	}
func main4() {
	defineMap()
	readWriteMap()
	checkExistenceWithCommaOk()
	usingMapAsSet()
	//usingMapAsSetWithStruct()
}

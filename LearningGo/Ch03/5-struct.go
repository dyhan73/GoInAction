package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func basicUsages() {
	var fred person
	bob := person{}

	fmt.Println(fred, bob)

	julia := person{
		"Julia",
		40,
		"cat",
	}
	fmt.Println(julia)
	julia = person{"Julia2", 41, "big cat"}
	fmt.Println(julia)

	beth := person{age: 30, pet: "dog"}
	fmt.Println(beth)
	fmt.Println(beth.pet)
	//beth = person{"Beth", 31, } // must fill all field without field names
}

func anonymousStruct() {
	var person struct {
		name string
		age  int
		pet  string
	}
	person.name = "bob"
	person.age = 50
	person.pet = "dog"
	fmt.Println(person)

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Println(pet)
}

func compareStruct() {
	type firstPerson struct {
		name string
		age  int
	}

	fp := firstPerson{"han", 40}
	fmt.Println(fp)

	type secondPerson struct {
		name string
		age  int
	}
	sp := secondPerson(fp)
	fmt.Println(sp)
	//if sp == fp {
	//	fmt.Println("sp is equal to fp")
	//}

	type thirdPerson struct {
		age  int
		name string
	}
	//tp := thirdPerson(fp)

	type fourthPerson struct {
		firstName string
		age       int
	}

	type fifthPerson struct {
		name          string
		age           int
		favoriteColor string
	}
	//fifp := fifthPerson(fp)

	var g struct {
		name string
		age  int
		//favoriteColor string
	}
	g = fp
	fmt.Println(g == fp)
}

func main() {
	basicUsages()
	anonymousStruct()
	compareStruct()
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Order struct {
	ID          string    `json:"id"`
	DateOrdered time.Time `json:"date_ordered"`
	CustomerId  string    `json:"customer_id"`
	Items       []Item    `json:"items"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func jsonEncodeDecode() {
	toFile := Person{
		Name: "Fred",
		Age:  40,
	}

	tmpFile, err := ioutil.TempFile(os.TempDir(), "sample-")
	if err != nil {
		panic(err)
	}
	fmt.Println("tmpFile.Name() :", tmpFile.Name())
	defer os.Remove(tmpFile.Name())
	err = json.NewEncoder(tmpFile).Encode(toFile)
	if err != nil {
		panic(err)
	}
	err = tmpFile.Close()
	if err != nil {
		panic(err)
	}

	tmpFile2, err := os.Open(tmpFile.Name())
	if err != nil {
		panic(err)
	}
	var fromFile Person
	err = json.NewDecoder(tmpFile2).Decode(&fromFile)
	if err != nil {
		panic(err)
	}
	err = tmpFile2.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", fromFile)
}

func main3() {
	jsonEncodeDecode()

	data := `{"name": "Fred", "age": 40}
{"name": "Mary", "age":21}
{"name": "Pat", "age":30}`
	dec := json.NewDecoder(strings.NewReader(data))
	var t Person
	for dec.More() {
		err := dec.Decode(&t)
		if err != nil {
			panic(err)
		}
		fmt.Println(t, t.Name, t.Age)
	}

	//var b bytes.Buffer
	//enc := json.NewEncoder(&b)
	//for _, input := range allInputs {
	//	t := process(input)

	//	err = enc.Encode(t)
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	//out := b.String()
	//fmt.Println(out)
}

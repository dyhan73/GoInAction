package main

import (
	"fmt"
	"sync"
)

type SlowComplicatedParser interface {
	Parse(string) string
}

var parser SlowComplicatedParser
var once sync.Once

func Parse(dataToParse string) string {
	once.Do(func() {
		parser = initParser()
	})
	return parser.Parse(dataToParse)
}

func initParser() SlowComplicatedParser {
	fmt.Println("initParser()")
	return MyParser{}
}

type MyParser struct{}

func (p MyParser) Parse(source string) string {
	return fmt.Sprintf("my parser's Parse() - %s", source)
}

func main511() {
	fmt.Println(Parse("asdfasdf"))
	fmt.Println(Parse("222222"))
	fmt.Println(Parse("333333"))
	fmt.Println(Parse("444444"))
}

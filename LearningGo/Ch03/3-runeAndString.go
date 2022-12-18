package main

import "fmt"

func main3() {
	var s string = "hellow there"
	var b byte = s[6]
	fmt.Println(s, b)

	s = "안녕하세요? 한대영입니다."
	var s2 string = s[3:6]

	fmt.Println(s, s2, len(s))

	s = "hello, 한"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)

}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"wordcount/words"
)

func main() {
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("파일을 열 때 오류가 발생했습니다.", err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)
	fmt.Printf("총 %d 개의 단어를 발견했습니다.\n", count)
}

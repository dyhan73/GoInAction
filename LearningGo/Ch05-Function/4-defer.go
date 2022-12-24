package main

import (
	"io"
	"log"
	"os"
)

func main4() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}

	//deferExam()

	// 자원을 정리하는 클로저를 반환하는 일반적인 패턴 예제
	f, closer, err := getFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	data := make([]byte, 200)
	for {
		count, err := f.Read(data)
		//fmt.Println("Read file : ", count)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}

func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		file.Close()
	}, err
}

func deferExam() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 200)
	for {
		count, err := f.Read(data)
		//fmt.Println("Read file : ", count)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}

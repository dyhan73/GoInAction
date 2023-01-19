package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}

	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func runCountLetters() {
	f, err := os.Open("LearningGo/Ch11-StandardLibrary/1-io.go")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	m, err := countLetters(f)
	if err != nil {
		fmt.Println("error received:", err)
		return
	}
	fmt.Println(m)
}

func examStringsNewReader() {
	s := "The quick brown fox jumped over the lazy dog"
	sr := strings.NewReader(s)
	counts, err := countLetters(sr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(counts)
}

func buildGZipReader(fileName string) (*gzip.Reader, func(), error) {
	r, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}
func examBuildGzipReader() error {
	r, closer, err := buildGZipReader("LearningGo/Ch11-StandardLibrary/asdf.zip")
	if err != nil {
		return err
	}
	defer closer()
	counts, err := countLetters(r)
	if err != nil {
		return err
	}
	fmt.Println(counts)
	return nil
}

func main1() {

	//runCountLetters()
	//examStringsNewReader()
	examBuildGzipReader()
}

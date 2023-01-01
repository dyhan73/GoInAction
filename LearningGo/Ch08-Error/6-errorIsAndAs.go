package main

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker2(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

func main6() {
	err := fileChecker2("not_here.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("That file doesn't exist")
		}
	}
}

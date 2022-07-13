package main

import (
	"log"
	"os"
)

// init 함수는 main 보다 먼저 호출
func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("Sherlock Holmes")
}

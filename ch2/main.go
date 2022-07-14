package main

import (
	_ "GoInAction/ch2/matchers" // matchers 패키지가 사용되지는 않으나 matchers 내부 코드의 init() 을 main() 전 실행하기 위해 _ 등록
	"GoInAction/ch2/search"
	"log"
	"os"
)

// init 함수는 main 보다 먼저 호출
func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("Russia")
}

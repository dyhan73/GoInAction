package example

import (
	"fmt"
)

type duration int

func (d duration) pretty() string {
	return fmt.Sprintf("기간 : %d", d)
}

func Main46() {
	fmt.Println(duration(42).pretty())
}

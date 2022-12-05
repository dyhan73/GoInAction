package example

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func Main35() {
	var b bytes.Buffer

	b.Write([]byte("안녕하세요?"))
	fmt.Fprintf(&b, "Go in action!\n")
	io.Copy(os.Stdout, &b)
}

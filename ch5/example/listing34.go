// io.Reader 와 io.Writer 인터페이스를 이용하여
// curl 을 간략히 작성

package example

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {
	//if len(os.Args) != 2 {
	//	fmt.Println("사용법: ./example2 <url>")
	//	os.Exit(-1)
	//}
}

func Main34() {
	//r, err := http.Get(os.Args[1])
	r, err := http.Get("http://naver.com")

	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, r.Body)
	fmt.Println(r.Header)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}

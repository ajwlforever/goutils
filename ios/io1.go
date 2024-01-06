package ios

import (
	"fmt"
	"io"
)

func Io1() {
	var reader io.Reader
	var p []byte
	n, err := reader.Read(p)
	if err != nil {
		return
	}
	fmt.Printf("n %d", n)
}

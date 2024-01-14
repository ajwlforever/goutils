package main

import (
	"bufio"
	"errors"
	"fmt"
)

type Writer1 struct {
}

type Writer2 struct {
}

func (w *Writer1) Write(p []byte) (n int, err error) {
	fmt.Println("writing: ", p)
	return 0, errors.New(fmt.Sprintf("boom!"))
}

func main() {
	w := new(Writer1)
	bw := bufio.NewWriterSize(w, 3)
	fmt.Println(bw.Buffered())
	bw.Write([]byte{'a'})
	fmt.Println(bw.Buffered())
	bw.Write([]byte{'b'})
	fmt.Println(bw.Buffered())
	bw.Write([]byte{'c'})
	fmt.Println(bw.Buffered())
	bw.Write([]byte{'d'})
	fmt.Println(bw.Buffered())

}

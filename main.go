package main

import (
	"fmt"

	"github.com/ajwlforever/goutils/hash"
)

func main() {
	mad5 := hash.StringMd5("ssss")
	fmt.Println(mad5)
}

package bingfa

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func test() {
	fmt.Println("sss")
}

// get string md5
func StringMd5(s string) string {
	md5 := md5.New()
	md5.Write([]byte(s))
	return hex.EncodeToString(md5.Sum(nil))
}

func Test_chan1() {
	test_chan1("xss")
}
func test_chan1(lll string) string {
	// Create
	pipeline := make(chan byte, 10)
	msg := 'c'
	go func() {
		for {
			// send msg to pipeline for all time
			pipeline <- byte(msg)
		}
	}()

	for {
		message := <-pipeline
		fmt.Printf("%s\n", message)
		fmt.Println(message)
	}

}

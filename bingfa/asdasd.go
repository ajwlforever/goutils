package bingfa

import (
	"fmt"
)

func test_chan1() {
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

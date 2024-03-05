package main

import "time"

var limit = make(chan int)
var work = []func(){
	func() { println("1"); time.Sleep(1 * time.Second); limit <- 1 },
	func() { println("2"); time.Sleep(1 * time.Second); limit <- 1 },
	func() { println("3"); time.Sleep(1 * time.Second); limit <- 1 },
	func() { println("4"); time.Sleep(1 * time.Second); limit <- 1 },
	func() { println("5"); time.Sleep(1 * time.Second); limit <- 1 },
}

// 对于无缓冲信道，当 <-done 执行时，必然要求 done <- 1 也已经执行。
func main() {
	for _, w := range work {
		go func(w func()) {
			w()
		}(w)
		<-limit
	}
	select {}
}

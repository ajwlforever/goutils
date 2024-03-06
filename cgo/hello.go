// hello.go
package main

/*
#include <stdio.h>

static void SayHello(const char* s) {
	puts(s);
}
*/
import "C"
import "fmt"

func main() {

	i := make(chan int, 2)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
		i <- 1
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}

	}()
	for i := 0; i < 1000; i++ {
		fmt.Println("ssss")
	}

	<-i
	<-i
}

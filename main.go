package main

import "fmt"

type F interface {
	f()
}
type s1 struct{}

func (s s1) f() {}

type s2 struct {
}

func (s *s2) f() {
}

func main() {
	// network.PollURL()
	//network.FetchURL("https://www.jb51.net/softs/708254.html")
	//
	// network.StartWeb()
	done := make(chan bool)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
		done <- true
	}()
	<-done
	fmt.Println("sss")

}

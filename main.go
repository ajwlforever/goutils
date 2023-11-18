package main

import (
	"github.com/ajwlforever/goutils/network"
	"time"
)

func NetWork1() {
	go network.StartServer()
	network.StartClient()

	time.Sleep(10000)

}

func main() {
	// network.PollURL()
	//network.FetchURL("https://www.jb51.net/softs/708254.html")
	network.StartWeb()
}

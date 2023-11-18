package network

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")

	var urlPath string = req.URL.Path
	tags := strings.Split(urlPath[1:], "/")
	if len(tags) == 2 {
		hello := tags[0]
		name := tags[1]
		if strings.EqualFold(hello, "hello") {
			fmt.Fprintf(w, "Hello,"+strings.ToUpper(name))
		}
	}

}

func HandleHTTP() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("localhost:8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

package network

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var urls = []string{
	"http://www.google.com/",
	"https://learnku.com/docs/the-way-to-go/153-access-and-read-pages/3705",
	"http://blog.golang.org/",
}

func PollURL() {
	// Execute an HTTP HEAD request for all url's
	// and returns the HTTP status string or an error string.
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
		}
		fmt.Println(url, ": ", resp.Status)
	}
}

func FetchURL(url string) (string, error) {
	res, err := http.Get(url)
	checkError(err)
	data, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Println("Got:", string(data))
	return string(data), err

}

func checkError(err error) {
	if err != nil {
		log.Fatal("Get: %v error", err)
	}
}

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

var message string

func init() {
	flag.StringVar(&message, "message", "Hello", "The Message to Log and Display")
	flag.Parse()
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	display := fmt.Sprintf("<h1>%s</h1><p>I'm a Docker app living on Diego!</p>", message)
	w.Write([]byte(display))
}

func instanceIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fetchIndex()))
}

func main() {
	go func() {
		for {
			fmt.Println(fmt.Sprintf("Diego Edge Docker App. Says %s (Instance %s)", message, fetchIndex()))
			time.Sleep(1 * time.Second)
		}
	}()

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/instance-index", instanceIndexHandler)

	http.ListenAndServe(":8080", nil)
}

func fetchIndex() string {
	index := os.Getenv("CF_INSTANCE_INDEX")
	if index == "" {
		index = os.Getenv("INSTANCE_INDEX")
	}
	return index
}

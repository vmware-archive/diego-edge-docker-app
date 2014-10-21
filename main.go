package main

import (
	"fmt"
	"time"
	"os"
	"net/http"
)

var instanceIndex string = os.Getenv("CF_INSTANCE_INDEX")

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello world</h1><p>I'm a Docker app living on Diego!</p>"))
}

func instanceIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(instanceIndex))
}

func main () {
	go func() {
		for {
			fmt.Println(fmt.Sprintf("Diego Edge Docker App. Says Hello (Instance %s)", instanceIndex))
			time.Sleep(1 * time.Second)
		}
	}()

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/instance-index", instanceIndexHandler)

	http.ListenAndServe(":8080", nil)
}

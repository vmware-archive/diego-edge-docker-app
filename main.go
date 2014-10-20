package main

import (
	"net/http"
)

func main () {
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello world</h1><p>I'm a Docker app living on Diego!</p>"))
	}

	http.HandleFunc("/", helloHandler)

	http.ListenAndServe(":8080", nil)
}

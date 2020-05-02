package main

import (
	"go-future/internal/app/hello_http"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello_http.HelloServer)
	http.HandleFunc("/time", hello_http.TimeServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}



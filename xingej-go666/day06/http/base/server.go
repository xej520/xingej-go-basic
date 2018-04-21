package main

import (
	"net/http"
	"io"
	"log"
)

func main() {
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":12345", nil))
	log.Println("held")
}

func HelloServer(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "hello, world!\n")
}



package main

import (
	"net/http"
	"net/http/httputil"
	"fmt"
)

func main() {
	resp, err := http.Get("http://localhost:12345/hello")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)

	if err!=nil {
		panic(err)
	}

	fmt.Printf("%s\n", s)

}





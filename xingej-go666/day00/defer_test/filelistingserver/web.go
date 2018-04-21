package main

import (
	"net/http"
	"xingej-go666/day00/defer_test/filelistingserver/filelisting"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(w http.ResponseWriter, r *http.Request){

}


func main() {
	http.HandleFunc("/list/", filelisting.HandleFileList)

	error := http.ListenAndServe(":8888", nil)
	if error != nil {
		panic(error)
	}

}



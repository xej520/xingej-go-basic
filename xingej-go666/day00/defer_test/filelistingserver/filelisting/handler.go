package filelisting

import (
	"net/http"
	"os"
	"io/ioutil"
)

func HandleFileList (writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path[len("/list/"):]  //  /list/fib.txt
	file, error := os.Open(path)
	if error != nil {
		panic(error)
	}

	defer file.Close()

	all, err := ioutil.ReadAll(file)

	if err != nil {
		panic(err)
	}

	writer.Write(all)
}

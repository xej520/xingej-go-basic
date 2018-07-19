package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	sparkSubmitPath := filepath.Join("/usr/local/spark", "bin", "/spark-submit")

	fmt.Println("SparkSubmitPath:\t", sparkSubmitPath)

}

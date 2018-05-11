//测试读取yaml文件的
package main

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
)

func main() {
	file, err := yaml.ReadFile("E:\\Program\\go2\\goPath\\src\\xingej-go\\xingej-go\\xingej-go666\\lib\\yaml\\nginx")

	if err != nil {
		panic(err.Error())
	}

	apiVersion, error := file.Get("apiVersion")
	if error != nil {
		panic(error.Error())
	}

	fmt.Println("=apiVersion===:\t", apiVersion)

}

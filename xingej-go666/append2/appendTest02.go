package main

import "fmt"

func main() {
	name := []string{"spark", "hadoop", "baidu", "docker", "k8s"}
	var nameList []string
	for key, value := range name {
		fmt.Printf("%s", key)
		env := fmt.Sprintf("%s%s=%s", "kafkaPreFix", key, value)
		nameList = append(nameList, env)
	}

	for key, value := range nameList {
		fmt.Println("key:\t", key, "\t", "value:\t", value)
	}

}

package main

import (
	"log"
	"os/exec"
)

func main() {

	command := "echo hello"
	cmd := exec.Command("/bin/bash", "-c", command)
	bytes, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}
	resp := string(bytes)
	log.Println(resp)

}

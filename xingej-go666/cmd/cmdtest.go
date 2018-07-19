package main

import (
	"os/exec"
)

func main() {

	lsCommand := exec.Command("tree")

	//lsCommand.Stdout = os.Stdout

	lsCommand.Run()
	//lsCommand.Start()
	//fmt.Println()

}

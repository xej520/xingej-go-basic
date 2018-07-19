package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd1 := "kubectl exec -it " + "ftp-testxjp3-worker-fn46n" + " mkdir /ftp" + " -n " + "testbonc"

	command := exec.Command("sh", "-c", cmd1)
	fmt.Println("------2.1------user---status:\t")

	if _, err := command.Output(); err != nil {
		fmt.Println("-----==================\n", err)
		//os.Exit(1)
	}

	command.Run()
	fmt.Println("------2.2------user---status:\t")
	command.Start()
	fmt.Println("------2.3------user---status:\t")

}

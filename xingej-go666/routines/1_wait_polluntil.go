package main

import (
	"k8s.io/apimachinery/pkg/util/wait"
	"time"
)

func main() {
	stop := make(chan struct{})
	wait.PollUntil(time.Second, func() (done bool, err error) {

		return false, nil
	}, stop)

	time.Sleep(time.Second * 5)
	stop <- struct{}{}

}

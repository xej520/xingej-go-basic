package main

import (
	"fmt"
	"regexp"
)

const text_3 = "my email is k8sAndDocker@google.com"

func main() {
	re := regexp.MustCompile(`.+@.+\..+`)
	match := re.FindString(text_3)

	fmt.Println(match)
}

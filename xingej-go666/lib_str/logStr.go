//包 str 测试
package main

import (
	"fmt"
	"github.com/mgutz/str"
)

func main() {
	logStr := "hello *s** b /dig?halolo/http"

	num := str.IndexOf(logStr, "/dig?", 0)

	result := str.Substr(logStr, num, 2)

	fmt.Println(result)
}

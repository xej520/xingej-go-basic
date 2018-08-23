package flag

import (
	"flag"
	"fmt"
	"testing"
)

func Test_flag_string(t *testing.T) {
	//命令行参数解析
	// 如果允许时，指定了   -clusterName="tianjin"
	//beijing 是默认值
	var name = flag.String("clusterName", "beijing", "the name of cluster")

	flag.Parse()

	fmt.Println(*name)
}

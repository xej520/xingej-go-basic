package main

import (
	"fmt"
	"strings"
)

func main() {
	var par map[string]string
	par = make(map[string]string)

	par["zo"] = "spark"
	par["z1"] = "spark1"
	par["z2"] = "spark2"
	par["z3"] = "spark3"

	result := UpdateUserConfig("", par)

	fmt.Println("------>\n", result)

}

// 更新用户参数
func UpdateUserConfig(oldParametes string, parameters map[string]string) (ftpcnf string) {
	//1、将旧参数，按照换行符进行分割
	oldparas := strings.Split(oldParametes, "\n")
	fmt.Println("--->\n", oldparas)
	fmt.Println("--2--->\n", len(oldparas))
	fmt.Println("--3--->\n", oldparas[0])
	var oldparasMap map[string]string
	oldparasMap = make(map[string]string)

	// 2、将旧参数(这里的旧参数，包括所有用户的配置，其实有点冗余了)，转换map类型

	if len(oldparas) > 1 {
		for _, elem := range oldparas {
			split := strings.Split(elem, "=")
			oldparasMap[split[0]] = split[1]
		}
	}

	// 3、将新传进来的参数，设置到旧参数里
	for key, value := range parameters {
		oldparasMap[key] = value
	}

	// 4、将旧参数由Map类型，转换成数组形式
	var configDate []string
	for k, v := range oldparasMap {
		// 对每一对kv参数，添加换行符
		configDate = append(configDate, k+"="+v+"\n")
	}

	// 5、转换成string类型，返回
	return strings.Join(configDate, "")
}

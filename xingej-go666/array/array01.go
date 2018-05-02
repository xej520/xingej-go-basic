package main

import "fmt"

type resource struct {
	url    string
	target string
	start  int
	end    int
}

func main() {

	var res []resource

	r1 := resource{
		url:    "http://localhost:8888",
		target: "394",
		start:  1,
		end:    98,
	}

	r2 := resource{
		url:    "http://localhost:8888",
		target: "394",
		start:  1,
		end:    98,
	}

	res = append(append(res, r1), r2)

	for _, key := range res {
		fmt.Println("====>:\t", key)
	}

}

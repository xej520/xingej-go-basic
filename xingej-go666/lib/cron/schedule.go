package cron

import (
	"fmt"
	"github.com/robfig/cron"
)

func main() {
	fmt.Println("--------------1-------------")
	schedule, err := cron.ParseStandard("@every 5m")
	fmt.Println("--------------2-------------")

	if err != nil {
		panic("==============")
		fmt.Println("=======================>:\n", err)
	}

	//schedule.Next(time.Time{})
	fmt.Println("----->:\t", schedule)

}

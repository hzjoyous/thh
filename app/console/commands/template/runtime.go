package template

import (
	"fmt"
	"time"
)

func init() {
	addConsole("runTime", "runTime",
		func() {
			fmt.Println("runTime")
			t := time.Now()
			fmt.Println("hello")
			time.Sleep(time.Second * 3)
			elapsed := time.Since(t)
			fmt.Println("app run time", elapsed)
		})
}

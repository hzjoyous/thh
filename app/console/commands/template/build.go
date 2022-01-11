package template

import (
	"fmt"
	"time"
)

func init() {
	addConsole("build", "build",
		func() {
			fmt.Println("build")
			times := "2020-09-18 15:04:05"
			tagStartTime, _ := time.Parse("2025-01-01 01:01:01", times)
			tagEndTime, _ := time.Parse("2025-12-31 01:01:01", times)

			fmt.Println(tagStartTime, tagEndTime)
		})
}

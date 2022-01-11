package template

import (
	"fmt"
)

func init() {
	addConsole("wechaty", "wechaty",
		func() {
			fmt.Println("wechaty")

		})
}

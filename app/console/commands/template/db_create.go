package template

import (
	"fmt"
	"thh/app/models/user"
	"time"
)

func init() {
	addConsole("dbCreate", "dbCreate",
		func() {
			user := user.User{Username: "nihao", Email: time.Now().String() + "@test.com"}
			user.SetPassword("niahao")
			err := user.Create()
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("dbCreate")
		})
}

package demo

import (
	"fmt"
	"github.com/spf13/cobra"
	"thh/app/models/Users"
	"time"
)

func init() {
	appendCommand(&cobra.Command{Use: "dbCreate", Short: "dbCreate", Run: dbCreate})
}
func dbCreate(cmd *cobra.Command, args []string) {
	user := Users.User{Username: "nihao", Email: time.Now().String() + "@test.com"}
	user.SetPassword("niahao")
	err := user.Create()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("dbCreate")
}

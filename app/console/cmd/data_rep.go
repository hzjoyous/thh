package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/gorm/utils"
	"thh/app/models/dataRep"
	"time"
)

func init() {
	appendCommand(&cobra.Command{Use: "demo:sqliteDBKV", Short: "demo:sqliteDBKV", Run: sqliteDBKV})
}

func sqliteDBKV(cmd *cobra.Command, args []string) {
	for i := 1; i <= 100; i++ {
		_ = dataRep.Set(utils.ToString(i), utils.ToString(time.Now().String()))
	}

	for i := 1; i <= 100; i++ {
		val := dataRep.Get(utils.ToString(i))
		fmt.Print(val)
	}
	fmt.Println()
}

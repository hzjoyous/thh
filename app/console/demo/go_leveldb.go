package demo

import (
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/gorm/utils"
	"thh/app/models/dataRep"
	"thh/helpers/db"
	"time"
)

func init() {
	appendCommand(&cobra.Command{Use: "demo:goLevelDB", Short: "demo:goLevelDB", Run: demoGoLevelDB})
	appendCommand(&cobra.Command{Use: "demo:sqliteDBKV", Short: "demo:sqliteDBKV", Run: sqliteDBKV})
}

func sqliteDBKV(cmd *cobra.Command, args []string) {
	for i := 1; i <= 100; i++ {
		_ = dataRep.GetDataRepository().Set(utils.ToString(i), utils.ToString(time.Now().String()))
	}

	for i := 1; i <= 100; i++ {
		val := dataRep.GetDataRepository().Get(utils.ToString(i))
		fmt.Print(val)
	}
	fmt.Println()
}

func demoGoLevelDB(cmd *cobra.Command, args []string) {
	var (
		kv = db.KVDB()
	)
	for i := 1; i <= 100; i++ {
		kv.Set(utils.ToString(i), utils.ToString(time.Now().String()))
	}
	for i := 1; i <= 100; i++ {
		fmt.Print(kv.Get(utils.ToString(i)))
	}
	fmt.Println()
	for i := 1; i <= 100; i++ {
		fmt.Print(kv.Get(utils.ToString(i)))
	}
	for i := 1; i <= 100; i++ {
		kv.Delete(utils.ToString(i))
	}
	fmt.Println()
	fmt.Println("after delete")
	for i := 1; i <= 100; i++ {
		fmt.Print(kv.Get(utils.ToString(i)))
	}
	fmt.Println("end")

}

package demo

import (
	"fmt"
	"gorm.io/gorm/utils"
	"thh/app/models/dataRep"
	"thh/helpers/db"
	"time"
)

func init() {
	addConsole("demo:goLevelDB", "demo:goLevelDB", demoGoLevelDB)
	addConsole("demo:sqliteDBKV", "demo:sqliteDBKV", sqliteDBKV)
}

func sqliteDBKV() {
	for i := 1; i <= 100; i++ {
		_ = dataRep.GetDataRepository().Set(utils.ToString(i), utils.ToString(time.Now().String()))
	}

	for i := 1; i <= 100; i++ {
		val := dataRep.GetDataRepository().Get(utils.ToString(i))
		fmt.Print(val)
	}
	fmt.Println()
}

func demoGoLevelDB() {
	var (
		kv = DB.KVDB()
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

package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"runtime"
	"thh/app/console"
	"thh/bootstrap"
	"thh/helpers/config"
)

var (
	app = bootstrap.Initialize()
)

/*
top		 main
up       http console job
up       controllers console
up       bootstrap
up       conf base
base     package helpers
*/
func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("Thousand-hand:start")
	fmt.Printf("Thousand-hand:useMem %d KB\n", m.Alloc/1024/8)
	switch len(os.Args) {
	case 1:
		app.ServerRun()
		break
	default:
		console.Execute()
		break
	}
	fmt.Println(config.GetString("app"))
}

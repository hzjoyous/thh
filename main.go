package main

import (
	_ "net/http/pprof"
	"thh/app/console"
	"thh/bootstrap"
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
	bootstrap.Initialize()
	console.Execute()

}

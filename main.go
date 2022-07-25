package main

import (
	"embed"
	_ "net/http/pprof"
	"thh/app/console"
	"thh/arms/app"
)

//go:embed web/*
var webFS embed.FS

//go:embed actor/dist/*
var actorFS embed.FS

//go:embed .env.example
var envExample string

/*
top		 main
up       http console job
up       controllers console
up       bootstrap
up       conf base
base     package helpers
*/
func main() {
	// 注册静态资源
	app.InitStart()
	app.WebRepSave(webFS)
	app.ActorSave(actorFS)
	app.EnvExample(envExample)

	console.Execute()
}

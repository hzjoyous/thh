package app

import "embed"

var webFS embed.FS
var actorFS embed.FS
var envExample string

func WebRepSave(dataWebFS embed.FS) {
	webFS = dataWebFS
}

func GetWebFS() embed.FS {
	return webFS
}

func ActorSave(dataWebFS embed.FS) {
	actorFS = dataWebFS
}

func GetActorFS() embed.FS {
	return actorFS
}

func EnvExample(data string) {
	envExample = data
}

func GetEnvExample() string {
	return envExample
}

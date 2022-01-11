package origin

import (
	"thh/base"
)

var commandList = make(map[string]base.Console)

func GetAllConsoles() map[string]base.Console {
	return commandList
}

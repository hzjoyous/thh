package example

import "thh/base"

var commandList = make(map[string]base.Console)

func GetAllConsoles() map[string]base.Console {
	return commandList
}

func addConsole(signature string, description string, handle func()) {
	c := base.Console{Signature: signature, Description: description, Handle: handle}
	commandList[c.Signature] = c
}

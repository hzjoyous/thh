package console

import (
	"fmt"
	"thh/base"
)

func getAllCommand() map[string]base.Console {
	return commandContainer
}

func addCommand(name string, desc string, handle func()) {
	c := base.Console{Signature: name, Description: desc, Handle: handle}
	commandContainer[c.Signature] = c
}

func p(newMap map[string]base.Console) {
	pushCommandList(newMap)
}

func pushCommandList(newMap map[string]base.Console) {
	commandContainer = consoleMapMerge(commandContainer, newMap)
}

func consoleMapMerge(mapA map[string]base.Console, mapB map[string]base.Console) map[string]base.Console {
	for consoleName, consoleEntity := range mapB {
		mapA[consoleName] = consoleEntity
	}
	return mapA
}

func Run(args []string) {
	fmt.Println("run")
	commandList := getAllCommand()
	var action string
	action = "help"
	if len(args) > 1 {
		action = args[1]
	}
	fmt.Println(action, ":")
	consoleEntity, ok := commandList[action]
	if ok {
		consoleEntity.Handle()
	} else {
		fmt.Println("command not found")
	}
}

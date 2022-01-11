package mirai_manager

import (
	"fmt"
	"log"
	"thh/app/service/remote_service"
	"thh/base"
)

var commandList = make(map[string]base.Console)

func GetAllConsoles() map[string]base.Console {
	return commandList
}
func addConsole(signature string, description string, handle func()) {
	c := base.Console{Signature: signature, Description: description, Handle: handle}
	commandList[c.Signature] = c
}

func init() {
	addConsole("z:mclient", "", MiraiClientManager)
	addConsole("z:goMclient", "", action)

}

func action() {
	c := remote_service.GoCqClientConnection("http://127.0.0.1:9091")
	a, e := c.SendGroupMsg()
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(a.String())
	}
}

func MiraiClientManager() {

	fmt.Println("start Mirai client ")

	client := remote_service.MiraiClientStd()

	result, err := client.About()
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}

	fmt.Println(result.String())
	r, _ := client.GroupList()
	fmt.Println(r.String())

	go scheduling(client)

	//quit := make(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	//fmt.Println("运行成功")
	//<-quit
	//fmt.Println("收到信号，结束")
}

package mirai

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"thh/app/service/remote_service"
)

func init() {
	appendCommand(&cobra.Command{Use: "z:mclient", Short: "", Run: MiraiClientManager})
	appendCommand(&cobra.Command{Use: "z:goMclient", Short: "", Run: action})

}

func action(cmd *cobra.Command, args []string) {
	c := remote_service.GoCqClientConnection("http://127.0.0.1:9091")
	a, e := c.SendGroupMsg()
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(a.String())
	}
}

func MiraiClientManager(cmd *cobra.Command, args []string) {

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
	client.AirFriendList(remote_service.GetTextMessage("过年好呀过年好~~~~~"))

	//go scheduling(client)

	//quit := gen(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	//fmt.Println("运行成功")
	//<-quit
	//fmt.Println("收到信号，结束")
}

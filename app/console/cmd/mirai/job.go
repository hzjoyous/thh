package mirai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cast"
	"thh/app/models/dataRep"
	"thh/app/service/remote_service"
	"time"
)

func scheduling(client *remote_service.MiraiHttpClient) {
	c := cron.New()
	var err error
	//upFunc(sendHello, client)()
	//_, err = c.AddFunc("30 8 * * *", upFunc(sendHello, client))
	_, err = c.AddFunc("0 7 * * *", upFunc(send622611442, client))
	_, err = c.AddFunc("*/15 9-22 * * *", upFunc(sendHello, client))
	//_, err = c.AddFunc("@every 20m", upFunc(continueSession, client))
	if err != nil {
		fmt.Println(err)
	}
	c.Run()
	fmt.Println("scheduling")
}

// sendHello
func send622611442(client *remote_service.MiraiHttpClient) {
	now := time.Now()
	str := bytes.Buffer{}
	str.WriteString("今天是")
	str.WriteString(now.Format("2006"))
	str.WriteString("年的，第")
	str.WriteString(cast.ToString(now.YearDay()))
	str.WriteString("天")
	fmt.Println(str.String())
	r, err := client.SendGroupMessage("622611442", str.String())
	if r != nil {
		fmt.Println(string(r.Body()), err)
	} else {
		fmt.Println(err)
	}
}
func sendHello(client *remote_service.MiraiHttpClient) {
	oldTime := dataRep.Get("sendHello")
	newTime := cast.ToString(time.Now().Format("20060102"))
	if oldTime < newTime {
		dataRep.Set("sendHello", newTime)
	}

	str := bytes.Buffer{}

	hiClient := remote_service.HiTokotoClientConnection("")
	resp, _ := hiClient.GetOneTokoto()
	var hitResp remote_service.HiTokotoResponse
	_ = json.Unmarshal(resp.Body(), &hitResp)
	str.WriteString(hitResp.Hitokoto)
	fmt.Println(str.String())

	r, err := client.SendGroupMessage("820744878", str.String())

	if r != nil {
		fmt.Println(string(r.Body()), err)
	} else {
		fmt.Println(err)
	}
}

func upFunc(f func(client *remote_service.MiraiHttpClient), client *remote_service.MiraiHttpClient) func() {
	return func() {
		f(client)
	}
}

func continueSession(client *remote_service.MiraiHttpClient) {
	result, err := client.Verify(client.GetAdminQQNumber())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("session 更新成功", result.String())
	}
}

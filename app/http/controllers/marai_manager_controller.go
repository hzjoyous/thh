package controllers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"thh/app/service/remote_service"
	"thh/helpers"
	Logger "thh/helpers/logger"
)

type request4MaraiManager struct {
	Type         string `json:"type"`
	MessageChain []struct {
		Type string `json:"type"`
		ID   int    `json:"id,omitempty"`
		Time int    `json:"time,omitempty"`
		Text string `json:"text,omitempty"`
	} `json:"messageChain"`
	Sender struct {
		ID                 int    `json:"id"`
		Nickname           string `json:"nickname"`
		Remark             string `json:"remark"`
		MemberName         string `json:"memberName"`
		SpecialTitle       string `json:"specialTitle"`
		Permission         string `json:"permission"`
		JoinTimestamp      int    `json:"joinTimestamp"`
		LastSpeakTimestamp int    `json:"lastSpeakTimestamp"`
		MuteTimeRemaining  int    `json:"muteTimeRemaining"`
		Group              struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			Permission string `json:"permission"`
		} `json:"group"`
	} `json:"sender"`
}

func MaraiManager(c *gin.Context) {
	rawData, _ := c.GetRawData()
	fmt.Println(string(rawData))
	Logger.Std().Info(string(rawData))
	// 重新写回，否则回出现EOF情况
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rawData))

	var requestData request4MaraiManager
	if err := c.BindJSON(&requestData); err != nil {
		fmt.Println(err)
		fmt.Println("检测失败")
	}
	client := remote_service.MiraiClientStd()
	requestMessage := ""
	switch requestData.Type {
	case "FriendMessage":
		for _, value := range requestData.MessageChain {
			if value.Type == "Plain" {
				requestMessage += value.Text
			}
		}
		if len(requestMessage) == 0 {
			requestMessage = "hi"
		}
		r, _ := client.SendFriendMessage(helpers.ToString(requestData.Sender.ID), requestMessage)
		fmt.Println(r.String())
		r, _ = client.SendGroupMessageChain("1146399464",
			remote_service.GetAtMessage("2817736127"),
			remote_service.GetTextMessage(requestMessage),
		)
		r, _ = client.SendGroupMessageChain("260671135",
			remote_service.GetAtMessage("31792690"),
			remote_service.GetTextMessage(requestMessage),
		)
		break
	case "GroupMessage":
		break
	}
	c.String(http.StatusOK, "ok")
}

package remote_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"runtime"
	"strconv"
	"sync"
	"thh/app/models/dataRep"
	"thh/helpers"
	"thh/helpers/config"
)

var miraiHttpClientStd = &MiraiHttpClient{}

// MiraiClientStd 标准客户端
func MiraiClientStd() *MiraiHttpClient {
	miraiHttpClientStd.once.Do(func() {
		authKey := config.GetString("authKey")
		adminQQNumber := config.GetString("adminQQNumber")
		host := config.GetString("miraiServerHost")
		fmt.Println(authKey, adminQQNumber, host)
		if len(host) == 0 {
			host = "http://127.0.0.1:9091"
		}
		miraiHttpClientStd.data = struct {
			sessionKey        string
			sessionVerifyTime int64
			adminQQNumber     string
			authKey           string
		}{adminQQNumber: adminQQNumber, authKey: authKey}
		miraiHttpClientStd.httpClient = resty.New().SetBaseURL(host)
	})
	return miraiHttpClientStd
}

type MiraiHttpClient struct {
	once       sync.Once
	lock       sync.Locker
	httpClient *resty.Client
	data       struct {
		sessionKey        string
		sessionVerifyTime int64
		adminQQNumber     string
		authKey           string
	}
}

func (itself *MiraiHttpClient) getSessionKey() string {
	return itself.data.sessionKey
}
func (itself *MiraiHttpClient) getSessionVerifyTime() string {
	return itself.data.sessionKey
}
func (itself *MiraiHttpClient) GetAdminQQNumber() string {
	return itself.data.adminQQNumber
}
func (itself *MiraiHttpClient) GetAuthKey() string {
	return itself.data.authKey
}
func (itself *MiraiHttpClient) SetSessionKey(session string) {
	funcName, file, line, ok := runtime.Caller(0)
	if ok {
		fmt.Println("func name: " + runtime.FuncForPC(funcName).Name())
		fmt.Printf("file: %s, line: %d\n", file, line)
		fmt.Println(session)
	}
	itself.data.sessionKey = session
}

type BaseMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// SendPBody 发送 Post Body
func (itself *MiraiHttpClient) SendPBody(url string, body map[string]interface{}) *resty.Response {
	body["sessionKey"] = itself.getSessionKey()
	resp, _ := itself.httpClient.R().SetBody(body).Post(url)
	return resp
}

// SendGQuery 发送 GET Query
func (itself *MiraiHttpClient) SendGQuery(url string, data map[string]string) *resty.Response {
	data["sessionKey"] = itself.getSessionKey()
	resp, _ := itself.httpClient.R().SetQueryParams(data).Get(url)
	return resp
}

func (itself *MiraiHttpClient) VerifySession() error {
	var authR authResponse
	var err error
	err = nil
	//oldSession := models.GetDataRepository().Get("session")
	//
	//if len(oldSession) > 0 {
	//	err = json.Unmarshal([]byte(oldSession), &authR)
	//	if len(authR.Session) > 0 {
	//		itself.SetSessionKey(authR.Session)
	//	}
	//}

	result, _ := itself.Auth(itself.GetAuthKey())
	err = json.Unmarshal([]byte(result.String()), &authR)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if authR.Code == 1 {
		return errors.New(authR.Msg)
	}
	data, _ := json.Marshal(authR)

	dataRep.GetDataRepository().Set("session", helpers.ToString(data))

	itself.SetSessionKey(authR.Session)

	result, err = itself.Verify(itself.GetAdminQQNumber())

	if err != nil {
		fmt.Println(err)
		return err
	}

	var baseMsg BaseMsg
	err = json.Unmarshal(result.Body(), &baseMsg)
	if err != nil {
		return err
	}
	if baseMsg.Code != 0 {
		return errors.New(baseMsg.Msg)
	}
	return err
}

type AboutResponse struct {
	Code         int    `json:"code"`
	Msg          string `json:"msg"`
	ErrorMessage string `json:"errorMessage"`
	Data         struct {
		Version string `json:"version"`
	} `json:"dataRep"`
}

// About 版本
func (itself *MiraiHttpClient) About() (resp *resty.Response, err error) {
	return itself.httpClient.R().SetQueryParams(map[string]string{}).Get("/about")
}

type authResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Session string `json:"session"`
}

// Auth 获取session
func (itself *MiraiHttpClient) Auth(authKey string) (resp *resty.Response, err error) {
	return itself.httpClient.R().SetBody(map[string]string{
		"authKey": authKey,
	}).Post("/auth")
}

func (itself *MiraiHttpClient) Verify(qq string) (resp *resty.Response, err error) {
	return itself.httpClient.R().SetBody(map[string]string{
		"sessionKey": itself.getSessionKey(),
		"qq":         qq,
	}).Post("/verify")
}

// Release 释放session
func (itself *MiraiHttpClient) Release() (resp *resty.Response, err error) {
	return itself.SendGQuery("/release", map[string]string{}), nil
}

type releaseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// SendFriendMessage 发送消息给friend
func (itself *MiraiHttpClient) SendFriendMessage(qq string, message string) (resp *resty.Response, err error) {
	return itself.SendPBody("/sendFriendMessage", map[string]interface{}{
		"target": qq,
		"messageChain": []map[string]interface{}{
			GetTextMessage(message),
		},
	}), nil
}

// SendFriendMessageList 发送消息给friend
func (itself *MiraiHttpClient) SendFriendMessageList(qq string, messageChainList ...map[string]interface{}) (resp *resty.Response, err error) {
	return itself.SendPBody("/sendFriendMessage", map[string]interface{}{
		"target":       qq,
		"messageChain": messageChainList,
	}), nil
}

type SendFriendMessageResponse struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	MessageID int    `json:"messageId"`
}

// SendTempMessage @Unverified 发送临时分组信息
func (itself *MiraiHttpClient) SendTempMessage(qq int, group int, messageType string, text string) (resp *resty.Response, err error) {
	return itself.SendPBody("/sendTempMessage", map[string]interface{}{
		"qq":    qq,
		"group": group,
		"messageChain": map[string]interface{}{
			"type": messageType,
			"text": text,
		},
	}), nil
}

// GroupList 获取好友列表
func (itself *MiraiHttpClient) GroupList() (resp *resty.Response, err error) {
	return itself.SendGQuery("/groupList", map[string]string{}), nil
}

// SendGroupMessage 发送小小给群组
func (itself *MiraiHttpClient) SendGroupMessage(GroupId string, message string) (resp *resty.Response, err error) {
	return itself.SendPBody("/sendGroupMessage", map[string]interface{}{
		"target": GroupId,
		"messageChain": []map[string]interface{}{
			{"type": "Plain", "text": message},
			//{"type": "Plain", "text": message},
			//{"type": "Face", "faceId": message},
			//{"type": "Plain", "url": message},
		},
		//"messageChain": []struct {
		//	Type string `json:"type"`
		//	Text string `json:"text,omitempty"`
		//	URL  string `json:"url,omitempty"`
		//}{{Type: "Plain", Text: message}},
	}), nil
}
func (itself *MiraiHttpClient) SendGroupMessageChain(GroupId string, msg ...map[string]interface{}) (resp *resty.Response, err error) {
	return itself.SendPBody("/sendGroupMessage", map[string]interface{}{
		"target":       GroupId,
		"messageChain": msg,
		//"messageChain": []struct {
		//	Type string `json:"type"`
		//	Text string `json:"text,omitempty"`
		//	URL  string `json:"url,omitempty"`
		//}{{Type: "Plain", Text: message}},
	}), nil
}

type SendGroupMessageResponse struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	MessageID int    `json:"messageId"`
}

// FriendList 获取好友列表
func (itself *MiraiHttpClient) FriendList() (resp *resty.Response, err error) {
	return itself.SendGQuery("/friendList", map[string]string{}), nil
}

type qqFriendEntity struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Remark   string `json:"remark"`
}

func (itself qqFriendEntity) getQQNumber() string {
	return strconv.Itoa(itself.ID)
}

// AirFriendList 群发消息
func (itself *MiraiHttpClient) AirFriendList(messageChainList ...map[string]interface{}) {
	result, _ := itself.FriendList()
	var qqFriendList []qqFriendEntity
	_ = json.Unmarshal([]byte(result.String()), &qqFriendList)
	for _, qqEntity := range qqFriendList {
		fmt.Println(qqEntity.ID, qqEntity.Remark)
		fmt.Println(qqEntity.getQQNumber())
		response, _ := itself.SendFriendMessageList(qqEntity.getQQNumber(), messageChainList...)
		fmt.Println(response)
	}
}

// GetTextMessage 文本信息
func GetTextMessage(message string) map[string]interface{} {
	return map[string]interface{}{"type": "Plain", "text": message}
}

// GetFaceMessage  1~289
func GetFaceMessage(message string) map[string]interface{} {
	return map[string]interface{}{"type": "Face", "faceId": message}
}

// GetImageMessage 图片信息
func GetImageMessage(message string) map[string]interface{} {
	return map[string]interface{}{"type": "Image", "faceId": message}
}

// at 某人
func GetAtMessage(qq string) map[string]interface{} {
	return map[string]interface{}{
		"type":   "At",
		"target": qq,
		//"display": "@Mirai",
	}
}

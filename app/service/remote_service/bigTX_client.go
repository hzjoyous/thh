package remote_service

import (
	"sync"

	"github.com/go-resty/resty/v2"
)

type BigTXClient struct {
	httpClient *resty.Client
}

var BigTXClientIns BigTXClient
var BigTXClientInsOnce sync.Once

func NewBigTXClient() BigTXClient {
	BigTXClientInsOnce.Do(func() {
		BigTXClientIns = BigTXClient{
			httpClient: resty.New().SetBaseURL("http://39.105.134.68:81"),
			//httpClient: resty.New().SetBaseURL("http://localhost:90/index.php"),
		}
		//BigTXClientIns.httpClient.SetProxy("http://127.0.0.1:11000")
	})
	return BigTXClientIns
}

func (itself *BigTXClient) SendEx() (resp *resty.Response, err error) {
	return itself.httpClient.R().SetFormData(map[string]string{
		"submit":  "发送物品",
		"account": "123aweqwe1",
		"money":   "",
		"items":   "0",
		"num":     "10",
	}).Post("/gm/")
}

func (itself *BigTXClient) RegNew(username string, password string) (resp *resty.Response, err error) {
	return itself.httpClient.R().SetFormData(map[string]string{
		"username":  username,
		"password":  password,
		"password1": password,
	}).Post("/api.php?act=reg")
}

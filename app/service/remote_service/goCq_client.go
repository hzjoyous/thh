package remote_service

import (
	"github.com/go-resty/resty/v2"
	"sync"
)

type GoCqClient struct {
	once       sync.Once
	httpClient *resty.Client
}

var GoCqClientIns = &GoCqClient{}

func GoCqClientConnection(host string) *GoCqClient {
	GoCqClientIns.once.Do(func() {
		if len(host) == 0 {
			host = "http://localhost:5700"
		}
		GoCqClientIns.httpClient = resty.New().SetBaseURL(host)
	})
	return GoCqClientIns
}

func (itself *GoCqClient) R() *resty.Request {
	return itself.httpClient.R()
}

//localhost:5700/send_group_msg?group_id=820744878&message=123123123
func (itself *GoCqClient) SendGroupMsg() (*resty.Response, error) {
	return itself.R().SetQueryParams(map[string]string{
		"group_id": "820744878",
		"message":  "溜了溜了",
	}).Get("send_group_msg")
}

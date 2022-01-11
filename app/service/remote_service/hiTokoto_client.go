package remote_service

import (
	"github.com/go-resty/resty/v2"
	"sync"
)

type HiTokotoClient struct {
	baseUri    string
	host       string
	httpClient *resty.Client
}

var HiTokotoClientIns HiTokotoClient

var hiTokotoClientInsOnce sync.Once

func HiTokotoClientConnection(host string) HiTokotoClient {
	hiTokotoClientInsOnce.Do(func() {
		if len(host) == 0 {
			host = "https://v1.hitokoto.cn/"
		}
		HiTokotoClientIns = HiTokotoClient{
			httpClient: resty.New().SetBaseURL(host),
		}
	})
	return HiTokotoClientIns
}

func (itself *HiTokotoClient) GetOneTokoto() (resp *resty.Response, err error) {
	return itself.httpClient.R().Get("/")
}

type HiTokotoResponse struct {
	ID         int    `json:"id"`
	UUID       string `json:"uuid"`
	Hitokoto   string `json:"hitokoto"`
	Type       string `json:"type"`
	From       string `json:"from"`
	FromWho    string `json:"from_who"`
	Creator    string `json:"creator"`
	CreatorUID int    `json:"creator_uid"`
	Reviewer   int    `json:"reviewer"`
	CommitFrom string `json:"commit_from"`
	CreatedAt  string `json:"created_at"`
	Length     int    `json:"length"`
}

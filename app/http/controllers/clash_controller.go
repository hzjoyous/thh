package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const config = `port: 7890
socks-port: 7891
redir-port: 7892
allow-lan: false
mode: global
log-level: silent
external-controller: '0.0.0.0:9090'
proxies:
    - {name: Charles, type: http, server: #{host}, port: 8888}
proxy-groups:
    - {name: global, type: relay, proxies: [Charles]}
`

func GetClashConfig(c *gin.Context) {
	host := c.Request.URL.Query().Get("host")
	if len(host) == 0 {
		c.String(500, "无法使用")
		return
	}
	configRep := strings.ReplaceAll(config, "#{host}", host)
	c.String(http.StatusOK, configRep)

}

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetClashConfig(c *gin.Context) {
	config := `port: 7890
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
	host := c.Request.URL.Query().Get("host")
	if len(host) == 0 {
		c.String(500, "无法使用")
		return
	}

	configRep := strings.Replace(config, "#{host}", host, 1)

	//c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	//c.Header("Pragma", "no-cache")
	//c.Header("Expires", "0")
	//c.Header("Content-Disposition", "attachment; filename=Charles-conf.yaml")
	//c.Data(http.StatusOK, "text/yaml; charset=utf-8", []byte(configRep))
	c.String(http.StatusOK, configRep)

}

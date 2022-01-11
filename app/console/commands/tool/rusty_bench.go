package tool

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

func init() {

	addConsole("rustyBenchTest", "", func() {
		const host = `https://www.baidu.com/`
		counter := 0
		httpClient := resty.New().
			SetBaseURL(host).
			SetHeader("Accept", "application/json").
			SetTimeout(time.Second * 30)
		for i := 0; i < 20; i++ {
			go func() {
				for {
					_, _ = httpClient.R().Get("")
					counter += 1
				}
			}()
		}
		now := 0
		for {

			time.Sleep(time.Second)
			fmt.Println(counter - now)
			now = counter
		}
	})
}

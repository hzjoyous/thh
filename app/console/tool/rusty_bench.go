package tool

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	appendCommand(&cobra.Command{Use: "z:rustyBenchTest", Short: "this is a z:rustyBenchTest", Run: rustyBenchTest})
}

func rustyBenchTest(cmd *cobra.Command, args []string) {
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
}

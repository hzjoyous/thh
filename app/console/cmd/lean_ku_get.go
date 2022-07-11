package cmd

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

func init() {
	appendCommand(CmdLeanKuGet)
}

var CmdLeanKuGet = &cobra.Command{
	Use:   "lean_ku_get",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runLeanKuGet,
}

func runLeanKuGet(cmd *cobra.Command, args []string) {
	tUrl := ""
	if len(tUrl) == 0 {
		return
	}
	for i := 127; i <= 142; i++ {
		//data := bytes.Buffer{}
		doc, err := htmlquery.LoadURL(tUrl + cast.ToString(i))
		if ifErr(err) {
			return
		}
		list := htmlquery.Find(doc, `//span[@class="topic-title"]/text()`)
		for _, value := range list {
			contentStr := htmlquery.InnerText(value)
			contentStr = strings.TrimSpace(contentStr)
			if len(contentStr) == 0 {
				continue
			}
			fmt.Println(contentStr)
			//data.WriteString(contentStr)
			//data.WriteString("\n")
		}
		//helpers.FilePutContents(`./learnkuList.txt`,data.Bytes(),true)
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

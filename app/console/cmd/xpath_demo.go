package cmd

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "xpath_demo",
		Short: "HERE PUTS THE COMMAND DESCRIPTION",
		Run:   runXpathDemo,
	})
}

func xpathDemo() {
	doc, err := htmlquery.Parse(strings.NewReader("asdas"))
	if ifErr(err) {
		return
	}
	htmlquery.InnerText(doc)
}

func runXpathDemo(cmd *cobra.Command, args []string) {

	filePath := "/Users/thh/Desktop/tmpData/tmpData.html"
	doc, err := htmlquery.LoadDoc(filePath)
	if ifErr(err) {
		return
	}
	list := htmlquery.Find(doc, "//dl[@class=\"source truncate-by-height\"]")
	for _, value := range list {
		content := htmlquery.FindOne(value, ".//span[1]")
		fmt.Println(htmlquery.InnerText(content))
	}
}

func ifErr(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

package console

import (
	"fmt"
	"strings"
	"thh/app/console/commands/demo"
	"thh/app/console/commands/example"
	"thh/app/console/commands/mirai_manager"
	"thh/app/console/commands/one"
	"thh/app/console/commands/origin"
	"thh/app/console/commands/template"
	"thh/app/console/commands/tool"
	"thh/base"
	"thh/helpers"

	"github.com/spf13/cobra"
)

var commandContainer = make(map[string]base.Console)

func init() {
	p(template.GetAllConsoles())
	p(example.GetAllConsoles())
	p(mirai_manager.GetAllConsoles())
	p(origin.GetAllConsoles())
	p(tool.GetAllConsoles())
	p(demo.GetAllConsoles())
	p(one.GetAllConsoles())

	for _, value := range getAllCommand() {

		rootCmd.AddCommand(&cobra.Command{
			Use:   value.Signature,
			Short: value.Description,
			Long:  value.Description,
			Run: func(value func()) func(cmd *cobra.Command, args []string) {
				return func(cmd *cobra.Command, args []string) {
					value()
				}
			}(value.Handle),
			//Run: func(cmd *cobra.Command, args []string) {
			//	value.Handle()
			//	fmt.Println(value.Signature)
			//},
			// 上面的将会导致value.handle 真正执行的时候从固定地址取，无法绑定真正的函数
		})
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "make:command",
		Short: "value.Description",
		Long:  "value.Description",
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) < 1 {
				fmt.Println("未传递名称")
				return
			}

			temple := `package template

import (
	"fmt"
)

func init() {
	addConsole("#{name}", "#{name}",
		func() {
			fmt.Println("#{name}")
		})
}
`
			newCommand := args[0]
			temple = strings.Replace(temple, "#{name}", newCommand, 3)
			_ = helpers.FilePutContents("./app/console/commands/template/"+newCommand+".go", []byte(temple), false)
		},
	})
}

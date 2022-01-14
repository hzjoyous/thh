package console

import (
	"thh/app/console/cmd"
	"thh/app/console/demo"
	"thh/app/console/example"
	"thh/app/console/gen"
	"thh/app/console/mirai"
	"thh/app/console/one"
	"thh/app/console/origin"
	"thh/app/console/template"
	"thh/app/console/tool"
)

func init() {
	rootCmd.AddCommand(gen.CmdMake)
	rootCmd.AddCommand(template.GetCommands()...)
	rootCmd.AddCommand(tool.GetCommands()...)
	rootCmd.AddCommand(one.GetCommands()...)
	rootCmd.AddCommand(origin.GetCommands()...)
	rootCmd.AddCommand(demo.GetCommands()...)
	rootCmd.AddCommand(example.GetCommands()...)
	rootCmd.AddCommand(mirai.GetCommands()...)
	rootCmd.AddCommand(cmd.GetCommands()...)
	rootCmd.AddCommand(CmdServe)

}

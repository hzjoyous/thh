package console

import (
	"thh/app/console/all"
	"thh/app/console/cmd"
	"thh/app/console/cmd/boom"
	"thh/app/console/cmd/demo"
	"thh/app/console/cmd/mirai"
	"thh/app/console/cmd/origin"
	"thh/app/console/cmd/tspider"
	"thh/app/console/gen"
	"thh/app/console/gmodel"
	"thh/app/console/one"
	"thh/app/console/shadow"
	"thh/app/console/tool"
)

func init() {
	rootCmd.AddCommand(gen.CmdMake)
	rootCmd.AddCommand(CmdServe)
	rootCmd.AddCommand(demo.GetCommands()...)
	rootCmd.AddCommand(tool.GetCommands()...)
	rootCmd.AddCommand(one.GetCommands()...)
	rootCmd.AddCommand(origin.GetCommands()...)
	rootCmd.AddCommand(mirai.GetCommands()...)
	rootCmd.AddCommand(cmd.GetCommands()...)
	rootCmd.AddCommand(gmodel.GetCommands()...)
	rootCmd.AddCommand(shadow.GetCommands()...)
	rootCmd.AddCommand(tspider.GetCommands()...)
	rootCmd.AddCommand(all.GetCommands()...)
	rootCmd.AddCommand(boom.GetCommands()...)

}

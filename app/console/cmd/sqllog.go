package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"thh/app/models/ActivityConfig"
	"thh/arms/logger"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "sqllog",
		Short: "HERE PUTS THE COMMAND DESCRIPTION",
		Run:   runSqllog,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	})
}

func runSqllog(cmd *cobra.Command, args []string) {
	logger.Info("sadasd")
	a := ActivityConfig.ActivityConfig{}
	ActivityConfig.Create(&a)
	fmt.Println(a)
}

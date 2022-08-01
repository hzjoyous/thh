package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "gateway",
		Short: "HERE PUTS THE COMMAND DESCRIPTION",
		Run:   runGateway,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	})
}

/**
每一个成功的id连接有一个id
reg可以获取前缀
*/
func runGateway(cmd *cobra.Command, args []string) {

}

func gwReg() {

}

func gwGateWay() {

}

func gwBusinessWorker() {

}

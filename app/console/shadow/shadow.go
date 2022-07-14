package shadow

import (
	"github.com/spf13/cobra"
)

func init() {
	appendCommand(CmdShadow)
}

var CmdShadow = &cobra.Command{
	Use:   "shadow",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runShadow,
	//Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

var shadowBody func()

func setShadowBody(f func()) {
	shadowBody = f
}
func getShadowBody() func() {
	return shadowBody
}
func runShadow(cmd *cobra.Command, args []string) {
	getShadowBody()()
}

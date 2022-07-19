package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "stdin",
		Short: "HERE PUTS THE COMMAND DESCRIPTION",
		Run:   runStdin,
		Args:  cobra.ExactArgs(0), // 只允许且必须传 0 个参数
	})
}

func runStdin(cmd *cobra.Command, args []string) {

	reader := bufio.NewReader(os.Stdin)
	for {
		result, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		fmt.Println(cast.ToString(result))
	}
}

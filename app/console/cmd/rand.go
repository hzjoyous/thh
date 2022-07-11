package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"thh/arms"
)

func init() {
	appendCommand(CmdRand)
}

var CmdRand = &cobra.Command{
	Use:   "rand",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runRand,
}

func runRand(cmd *cobra.Command, args []string) {

	fmt.Println(arms.RandomString(10))
}

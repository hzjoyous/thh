package console

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"runtime"
	"thh/helpers"
)

// RegisterDefaultCmd 注册默认命令
func RegisterDefaultCmd(rootCmd *cobra.Command, subCmd *cobra.Command) {
	cmd, _, err := rootCmd.Find(os.Args[1:])
	firstArg := helpers.FirstElement(os.Args[1:])
	if err == nil && cmd.Use == rootCmd.Use && firstArg != "-h" && firstArg != "--help" {
		args := append([]string{subCmd.Use}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "thh",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRun: func(command *cobra.Command, args []string) {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Println("Thousand-hand:start")
		fmt.Printf("Thousand-hand:useMem %d KB\n", m.Alloc/1024/8)
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	RegisterDefaultCmd(rootCmd, CmdServe)
	cobra.CheckErr(rootCmd.Execute())
}


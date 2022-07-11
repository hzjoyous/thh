package console

import (
	"github.com/spf13/cobra"
	"os"
	"thh/arms"
	"thh/bootstrap"
)

// RegisterDefaultCmd 注册默认命令
func RegisterDefaultCmd(rootCmd *cobra.Command, subCmd *cobra.Command) {
	cmd, _, err := rootCmd.Find(os.Args[1:])
	firstArg := arms.FirstElement(os.Args[1:])
	if err == nil && cmd.Use == rootCmd.Use && firstArg != "-h" && firstArg != "--help" {
		args := append([]string{subCmd.Use}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "thh",
	Short: "A brief description of your application",
	Long:  `thh`,
	PersistentPreRun: func(command *cobra.Command, args []string) {
		bootstrap.Initialize()
		//var m runtime.MemStats
		//runtime.ReadMemStats(&m)
		//fmt.Println("Thousand-hand:start")
		//fmt.Printf("Thousand-hand:useMem %d KB\n", m.Alloc/1024/8)
	},
	//Run: cmd.RunBoom,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	//RegisterDefaultCmd(rootCmd, CmdServe)
	cobra.CheckErr(rootCmd.Execute())
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sync"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "syncmap",
		Short: "HERE PUTS THE COMMAND DESCRIPTION",
		Run:   runSyncmap,
		//Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	})
}

func runSyncmap(cmd *cobra.Command, args []string) {
	m := sync.Map{}
	m.Store(1, 1)
	v, ok := m.Load(1)
	if !ok {
		fmt.Println("加载失败")
	}
	if value, ok := v.(int); ok {
		fmt.Println(value)
	}
}

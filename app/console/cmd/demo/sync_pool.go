package demo

import (
	"fmt"
	"github.com/spf13/cobra"
	"sync"
)

func init() {
	appendCommand(&cobra.Command{Use: "syncPool", Short: "syncPool", Run: syncPool})
}
func syncPool(cmd *cobra.Command, args []string) {
	type cat struct {
		Name string
	}
	var catPool = sync.Pool{
		New: func() any {
			return &cat{}
		},
	}
	cat1 := catPool.Get()
	var catTmp *cat
	catTmp, isCat := cat1.(*cat)

	if !isCat {
		return
	}
	catTmp.Name = "小红"

	catPool.Put(catTmp)

	cat1 = catPool.Get()
	catTmp, isCat = cat1.(*cat)
	if !isCat {
		return
	}
	fmt.Println(catTmp.Name)
	// 使用后不放回pool

	cat1 = catPool.Get()
	catTmp, isCat = cat1.(*cat)
	if !isCat {
		return
	}
	fmt.Println(catTmp.Name)

	fmt.Println("syncPool")
}

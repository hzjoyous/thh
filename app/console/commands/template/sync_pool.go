package template

import (
	"fmt"
	"sync"
)

func init() {
	addConsole("syncPool", "syncPool",
		func() {
			type cat struct {
				Name string
			}
			var catPool = sync.Pool{
				New: func() interface{} {
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
		})
}

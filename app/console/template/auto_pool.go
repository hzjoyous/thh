package template

import (
	"fmt"
	"github.com/spf13/cobra"
	"sync"
)

func init() {
	appendCommand(&cobra.Command{Use: "autoPool", Short: "autoPool", Run: autoPool})
}

type dog struct {
}
type superPool struct {
	pool sync.Pool
}

var sp = superPool{
	pool: sync.Pool{
		New: func() interface{} {
			return &dog{}
		},
	},
}

func (itself *superPool) gp(dogUp func(dogEntity *dog)) {
	dogI := itself.pool.Get()
	dogEntity, ok := dogI.(*dog)
	if !ok {
		return
	}
	dogUp(dogEntity)

	itself.pool.Put(dogEntity)
}

func (itself *superPool) dogRead(dogEntity *dog) {
	fmt.Println("dog")
}

func autoPool(cmd *cobra.Command, args []string) {
	sp.gp(sp.dogRead)

	fmt.Println("autoPool")
}

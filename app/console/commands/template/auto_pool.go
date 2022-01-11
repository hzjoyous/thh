package template

import (
	"fmt"
	"sync"
)

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

func init() {
	addConsole("autoPool", "autoPool",
		func() {

			sp.gp(sp.dogRead)

			fmt.Println("autoPool")
		})
}

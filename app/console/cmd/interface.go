package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	appendCommand(&cobra.Command{Use: "p:Interface", Short: "", Run: practiceInterface})
}

type animal interface {
	eat() int
	getHand() *hand
}

type Dog struct {
	name string
	hand *hand
}

func (d *Dog) eat() int {
	fmt.Println("1")
	return 1
}

func (d *Dog) getHand() *hand {
	return d.hand
}

type hand struct {
}

func practiceInterface(cmd *cobra.Command, args []string) {

	var dog animal
	dog = &Dog{}
	dog.eat()
}

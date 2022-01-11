package example

import (
	"fmt"
)

func init() {
	addConsole("practiceInterface", "this is a reflectDemo", practiceInterface)
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

func practiceInterface() {

	var dog animal
	dog = &Dog{}
	dog.eat()
}

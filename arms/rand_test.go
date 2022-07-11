package arms

import (
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {
	str := RandomString(10)
	fmt.Println(str)
}

func TestIdMakerP(t *testing.T) {
	idm := IdMakerInOnP{}
	idm.SetStartId(1000)
	fmt.Println(idm.Get())
	fmt.Println(idm.Get())
	fmt.Println(idm.Get())
	fmt.Println(idm.Get())
	idm.SetStartId(1000)
	fmt.Println(idm.Get())
	fmt.Println(idm.Get())
	fmt.Println(idm.Get())

	counter := Counter{}
	go func() {
		fmt.Println(counter.Add(), "end")
	}()
	fmt.Println(counter.Get())
}

func TestPassword(t *testing.T) {
	tmpPassWord := "asdasdas"
	password := MakePassword(tmpPassWord)
	err := VerifyPassword(password, tmpPassWord)
	if err != nil {
		fmt.Println(err)
	}
}

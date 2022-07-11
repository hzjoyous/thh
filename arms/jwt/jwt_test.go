package jwt

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateNewToken(t *testing.T) {
	token, err := CreateNewToken(123456, time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(token)
	userId, err := VerifyToken(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(userId)
	fmt.Println("token还可以用")

	time.Sleep(time.Second * 2)
	fmt.Println(token)
	userId, err = VerifyToken(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(userId)
}

func TestVerifyTokenWithFresh(t *testing.T) {
	token, err := CreateNewToken(123456, 5)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("token created : ", token)
	userId, newToken, err := VerifyTokenWithFresh(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("token can use")
	if newToken != token {
		token = newToken
	}
	fmt.Println(userId)

	time.Sleep(time.Second * 2)
	fmt.Println(token)
	userId, err = VerifyToken(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(userId)
}

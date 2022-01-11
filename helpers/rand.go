package helpers

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandTrue(Molecular int, Denominator int) bool {
	return rand.Intn(Denominator) < Molecular
}

func RandomString(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	lettersLen := len(letters)
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(lettersLen)]
	}
	return string(b)
}

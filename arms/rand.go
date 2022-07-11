package arms

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandTrue(Molecular int, Denominator int) bool {
	return rand.Intn(Denominator) < Molecular
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var lettersLen = len(letters)

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(lettersLen)]
	}
	return string(b)
}

func RandomNum(n int) int {
	return rand.Intn(n)
}

// IdMakerInOnP idMaker
type IdMakerInOnP struct {
	id   uint64
	lock sync.Mutex
	once sync.Once
}

func (itself *IdMakerInOnP) SetStartId(id uint64) {
	itself.once.Do(func() {
		itself.lock.Lock()
		defer itself.lock.Unlock()
		itself.id = id
	})
}

func (itself *IdMakerInOnP) Get() uint64 {
	itself.lock.Lock()
	defer itself.lock.Unlock()
	itself.id += 1
	return itself.id
}

type Counter struct {
	lock   sync.Mutex
	number int64
}

func (itself *Counter) Add() int64 {
	itself.lock.Lock()
	defer itself.lock.Unlock()
	itself.number += 1
	return itself.number
}

func (itself *Counter) Get() int64 {
	itself.lock.Lock()
	defer itself.lock.Unlock()
	return itself.number
}

func MakePassword(password string) string {
	//生成16位 Salt
	salt := RandomString(16)
	//计算 Salt 和密码组合的SHA1摘要
	hash := sha1.New()
	_, _ = hash.Write([]byte(password + salt))
	bs := hex.EncodeToString(hash.Sum(nil))
	//存储 Salt 值和摘要， ":"分割
	return salt + ":" + bs
}

func VerifyPassword(SecretPassword, inputPassword string) error {
	passwordStore := strings.Split(SecretPassword, ":")
	if len(passwordStore) != 2 {
		return errors.New("no pass")
	}
	//计算 Salt 和密码组合的SHA1摘要
	hash := sha1.New()
	_, _ = hash.Write([]byte(inputPassword + passwordStore[0]))
	bs := hex.EncodeToString(hash.Sum(nil))
	if bs == passwordStore[1] {
		return nil
	}
	return errors.New("input password error")
}

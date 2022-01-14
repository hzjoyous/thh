package user

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"thh/app/models"
	"thh/helpers/db"
)

func init() {
	ur = new(userRepository)
}

// User 用户模型
type User struct {
	models.BaseModel

	Username string `gorm:"type:varchar(255);not null;default:'';"  json:"username"`
	Email    string `gorm:"type:varchar(255);index:idx_email,unique;default:'';"  json:"email"`
	Password string `gorm:"type:varchar(255);default:'';"  json:"password"`
}

func MakeUser(name string, password string, email string) *User {
	user := User{Username: name, Email: email}
	user.SetPassword(password)
	return &user
}

func randStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (entity *User) SetPassword(password string) *User {
	//生成16位 Salt
	salt := randStringRunes(16)
	//计算 Salt 和密码组合的SHA1摘要
	hash := sha1.New()
	_, _ = hash.Write([]byte(password + salt))
	bs := hex.EncodeToString(hash.Sum(nil))
	//存储 Salt 值和摘要， ":"分割
	entity.Password = salt + ":" + bs
	return entity
}

func (entity *User) Create() (err error) {
	if err = db.SqlDBIns().Create(&entity).Error; err != nil {
		return err
	}
	return nil
}

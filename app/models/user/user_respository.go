package user

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"strings"
	"thh/helpers/db"
)

func init() {
	ur = new(userRepository)
}

type userRepository struct {
	model User
}

var ur *userRepository

func (itself *userRepository) getModel() User {
	return itself.model
}

func (itself *userRepository) DB() {
	db.SqlDBIns().Model(itself.getModel())
}

func (itself *userRepository) Find() []User {
	var users []User
	db.SqlDBIns().Find(&users)
	return users
}

func (itself *userRepository) GetById(id uint64) (User, error) {
	var users User
	result := db.SqlDBIns().First(&users, id)
	return users, result.Error
}

func (itself *userRepository) Verify(username string, password string) (User, error) {
	var user User
	err := db.SqlDBIns().Where("name = ? ", username).Find(&user).Error
	if err != nil {
		return user, err
	}
	passwordStore := strings.Split(user.Password, ":")
	if len(passwordStore) != 2 {
		return user, errors.New("no pass")
	}
	//计算 Salt 和密码组合的SHA1摘要
	hash := sha1.New()
	_, _ = hash.Write([]byte(password + passwordStore[0]))
	bs := hex.EncodeToString(hash.Sum(nil))
	if err != nil {
		return user, err
	}
	if bs == passwordStore[1] {
		return user, nil
	}
	return user, err
}

func UserRepository() *userRepository {
	return ur
}

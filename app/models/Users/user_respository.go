package Users

import (
	"thh/arms"
	"thh/conf/dbconnect"
)

func GetBy(field, value string) (userOpLock User) {
	dbconnect.Std().Where("? = ?", field, value).First(&userOpLock)
	return
}

func GetById(id uint64) (User, error) {
	var users User
	result := dbconnect.Std().First(&users, id)
	return users, result.Error
}

func Verify(username string, password string) (User, error) {
	var user User
	err := dbconnect.Std().Where("name = ? ", username).Find(&user).Error
	if err != nil {
		return user, err
	}
	err = arms.VerifyPassword(user.Password, password)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

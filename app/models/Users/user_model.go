package Users

import (
	"thh/app/models"
	"thh/arms"
	"thh/conf/dbconnect"
)

// User 用户模型
type User struct {
	models.BaseModel

	Username string `gorm:"type:varchar(255);not null;default:'';"  json:"username"`
	Email    string `gorm:"type:varchar(255);index:idx_email,unique;default:'';"  json:"email"`
	Password string `gorm:"type:varchar(255);default:'';"  json:"password"`
}

func (User) TableName() string {
	return "user"
}

func MakeUser(name string, password string, email string) *User {
	user := User{Username: name, Email: email}
	user.SetPassword(password)
	return &user
}

func (itself *User) Create() (err error) {
	if err = dbconnect.Std().Create(&itself).Error; err != nil {
		return err
	}
	return nil
}

func (itself *User) Save() (rowsAffected int64) {
	result := dbconnect.Std().Save(&itself)
	return result.RowsAffected
}

func (itself *User) Delete() (rowsAffected int64) {
	result := dbconnect.Std().Delete(&itself)
	return result.RowsAffected
}

func (itself *User) SetPassword(password string) *User {
	itself.Password = arms.MakePassword(password)
	return itself
}

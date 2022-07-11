package Users

import (
	"time"
)

const tableName = "users"
const pid = "id"
const fieldCreatedAt = "created_at"
const fieldUpdatedAt = "updated_at"
const fieldDeletedAt = "deleted_at"
const fieldUsername = "username"
const fieldEmail = "email"
const fieldPassword = "password"

type Users struct {
	Id        uint64     `gorm:"primaryKey;column:id;type:bigint unsigned;not null;" json:"id"`          //
	CreatedAt *time.Time `gorm:"column:created_at;type:datetime(3);" json:"createdAt"`                   //
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime(3);" json:"updatedAt"`                   //
	DeletedAt *time.Time `gorm:"column:deleted_at;type:datetime(3);" json:"deletedAt"`                   //
	Username  string     `gorm:"column:username;type:varchar(255);not null;default:'';" json:"username"` //
	Email     string     `gorm:"column:email;type:varchar(255);not null;default:'';" json:"email"`       //
	Password  string     `gorm:"column:password;type:varchar(255);not null;default:'';" json:"password"` //

}

func (Users) TableName() string {
	return tableName
}

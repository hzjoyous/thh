package Role

import (
	"time"
)

const tableName = "role"
const pid = "id"
const fieldCode = "code"
const fieldName = "name"
const fieldValid = "valid"
const fieldCreatedAt = "created_at"
const fieldUpdatedAt = "updated_at"
const fieldDeletedAt = "deleted_at"

type Role struct {
	Id        uint64     `gorm:"primaryKey;column:id;autoIncrement;not null;" json:"id"`                                //
	Code      string     `gorm:"column:code;type:varchar(200);not null;default:'';" json:"code"`                        // 权限code
	Name      string     `gorm:"column:name;type:varchar(32);not null;default:'';" json:"name"`                         // 权限名
	Valid     uint8      `gorm:"column:valid;type:tinyint unsigned;not null;default:1;" json:"valid"`                   //
	CreatedAt time.Time  `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"createdAt"` //
	UpdatedAt time.Time  `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"updatedAt"` //
	DeletedAt *time.Time `gorm:"column:deleted_at;type:timestamp;" json:"deletedAt"`                                    //

}

// func (itself *Role) BeforeSave(tx *gorm.DB) (err error) {}
// func (itself *Role) BeforeCreate(tx *gorm.DB) (err error) {}
// func (itself *Role) AfterCreate(tx *gorm.DB) (err error) {}
// func (itself *Role) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (itself *Role) AfterUpdate(tx *gorm.DB) (err error) {}
// func (itself *Role) AfterSave(tx *gorm.DB) (err error) {}
// func (itself *Role) BeforeDelete(tx *gorm.DB) (err error) {}
// func (itself *Role) AfterDelete(tx *gorm.DB) (err error) {}
// func (itself *Role) AfterFind(tx *gorm.DB) (err error) {}

func (Role) TableName() string {
	return tableName
}

package Permission

import (
	"time"
)

const tableName = "permission"
const pid = "id"
const fieldCode = "code"
const fieldName = "name"
const fieldPid = "pid"
const fieldValid = "valid"
const fieldCreatedAt = "created_at"
const fieldUpdatedAt = "updated_at"
const fieldDeletedAt = "deleted_at"

type Permission struct {
	Id        uint64     `gorm:"primaryKey;column:id;autoIncrement;not null;" json:"id"`                                //
	Code      string     `gorm:"column:code;type:varchar(200);not null;default:'';" json:"code"`                        // 权限code
	Name      string     `gorm:"column:name;type:varchar(32);not null;default:'';" json:"name"`                         // 权限名
	Pid       uint       `gorm:"column:pid;type:int unsigned;not null;default:0;" json:"pid"`                           //
	Valid     uint8      `gorm:"column:valid;type:tinyint unsigned;not null;default:1;" json:"valid"`                   //
	CreatedAt time.Time  `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"createdAt"` //
	UpdatedAt time.Time  `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"updatedAt"` //
	DeletedAt *time.Time `gorm:"column:deleted_at;type:timestamp;" json:"deletedAt"`                                    //

}

// func (itself *Permission) BeforeSave(tx *gorm.DB) (err error) {}
// func (itself *Permission) BeforeCreate(tx *gorm.DB) (err error) {}
// func (itself *Permission) AfterCreate(tx *gorm.DB) (err error) {}
// func (itself *Permission) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (itself *Permission) AfterUpdate(tx *gorm.DB) (err error) {}
// func (itself *Permission) AfterSave(tx *gorm.DB) (err error) {}
// func (itself *Permission) BeforeDelete(tx *gorm.DB) (err error) {}
// func (itself *Permission) AfterDelete(tx *gorm.DB) (err error) {}
// func (itself *Permission) AfterFind(tx *gorm.DB) (err error) {}

func (Permission) TableName() string {
	return tableName
}

package RolePermission

import (
	"time"
)

const tableName = "role_permission"
const pid = "id"
const fieldRoleId = "role_id"
const fieldPermissionId = "permission_id"
const fieldMultiSelect = "multi_select"
const fieldCreatedAt = "created_at"
const fieldUpdatedAt = "updated_at"
const fieldDeletedAt = "deleted_at"

type RolePermission struct {
	Id           uint64     `gorm:"primaryKey;column:id;autoIncrement;not null;" json:"id"`                                //
	RoleId       uint64     `gorm:"column:role_id;type:bigint unsigned;not null;default:0;" json:"roleId"`                 //
	PermissionId uint64     `gorm:"column:permission_id;type:bigint unsigned;not null;default:0;" json:"permissionId"`     //
	MultiSelect  uint8      `gorm:"column:multi_select;type:tinyint unsigned;not null;default:1;" json:"multiSelect"`      //
	CreatedAt    time.Time  `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"createdAt"` //
	UpdatedAt    time.Time  `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"updatedAt"` //
	DeletedAt    *time.Time `gorm:"column:deleted_at;type:timestamp;" json:"deletedAt"`                                    //

}

// func (itself *RolePermission) BeforeSave(tx *gorm.DB) (err error) {}
// func (itself *RolePermission) BeforeCreate(tx *gorm.DB) (err error) {}
// func (itself *RolePermission) AfterCreate(tx *gorm.DB) (err error) {}
// func (itself *RolePermission) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (itself *RolePermission) AfterUpdate(tx *gorm.DB) (err error) {}
// func (itself *RolePermission) AfterSave(tx *gorm.DB) (err error) {}
// func (itself *RolePermission) BeforeDelete(tx *gorm.DB) (err error) {}
// func (itself *RolePermission) AfterDelete(tx *gorm.DB) (err error) {}
// func (itself *RolePermission) AfterFind(tx *gorm.DB) (err error) {}

func (RolePermission) TableName() string {
	return tableName
}

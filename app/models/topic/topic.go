package topic

import (
	"thh/app/models"
)

type Topic struct {
	models.BaseModel
	State   int8   `gorm:"type:tinyint(3);not null;default:0;"  json:"state"`
	Title   string `gorm:"type:varchar(255)"`
	Content string
}

func (Topic) TableName() string {
	return "message_board_topic"
}

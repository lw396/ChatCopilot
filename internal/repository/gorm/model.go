package gorm

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `gorm:"type:timestamp null;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"type:timestamp null;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type GroupContact struct {
	Model
	UsrName         string `gorm:"column:usr_name"`
	Nickname        string `gorm:"column:nickname"`
	HeadImgUrl      string `gorm:"column:head_img"`
	ChatRoomMemList string `gorm:"column:group_member"`
	DBName          string `gorm:"column:db_name"`
}

type MessageContent struct {
}

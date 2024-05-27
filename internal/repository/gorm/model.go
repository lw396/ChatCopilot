package gorm

import (
	"time"

	"github.com/lw396/WeComCopilot/internal/model"
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
	Status          uint8  `gorm:"column:status"`
}

type ContactPerson struct {
	Model
	UsrName    string `gorm:"column:usr_name"`
	Nickname   string `gorm:"column:nickname"`
	Remark     string `gorm:"column:remark"`
	HeadImgUrl string `gorm:"column:head_img_url"`
	Sex        int64  `gorm:"column:sex"`
	Type       int64  `gorm:"column:type"`
	DBName     string `gorm:"column:db_name"`
	Status     uint8  `gorm:"column:status"`
}

type MessageContent struct {
	LocalID     int64             `gorm:"primaryKey;column:local_id"`
	SvrID       int64             `gorm:"column:svr_id"`
	CreateTime  int64             `gorm:"column:create_time"`
	Content     string            `gorm:"column:content"`
	Status      int64             `gorm:"column:status"`
	ImgStatus   int64             `gorm:"column:img_status"`
	MessageType model.MessageType `gorm:"column:message_type"`
	Des         bool              `gorm:"column:des"`
	Source      string            `gorm:"column:source"`
	VoiceText   string            `gorm:"column:vice_text"`
	Seq         int64             `gorm:"column:seq"`
}

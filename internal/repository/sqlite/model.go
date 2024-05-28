package sqlite

import (
	"github.com/lw396/WeComCopilot/internal/model"
)

// Group
type GroupContact struct {
	UsrName           string `gorm:"column:m_nsUsrName"`
	Nickname          string `gorm:"column:nickname"`
	Type              int64  `gorm:"column:m_uiType"`
	ImgStatus         string `gorm:"column:m_nsImgStatus"`
	HeadImgUrl        string `gorm:"column:m_nsHeadImgUrl"`
	ChatRoomMemList   string `gorm:"column:m_nsChatRoomMemList"`
	ChatRoomAdminList string `gorm:"column:m_nsChatRoomAdminList"`
	ChatRoomStatus    int64  `gorm:"column:m_uiChatRoomStatus"`
	ChatRoomVersion   int64  `gorm:"column:m_uiChatRoomVersion"`
	ChatRoomType      int64  `gorm:"column:m_uiChatRoomType"`
	DBName            string `gorm:"column:db_name"`
}

func (GroupContact) TableName() string {
	return "GroupContact"
}

type ContactPerson struct {
	UsrName            string `gorm:"column:m_nsUsrName"`
	ConType            int64  `gorm:"column:m_uiConType"`
	Nickname           string `gorm:"column:nickname"`
	FullPingyin        string `gorm:"column:m_nsFullPY"`
	ShortPingyin       string `gorm:"column:m_nsShortPY"`
	Remark             string `gorm:"column:m_nsRemark"`
	RemarkFullPingyin  string `gorm:"column:m_nsRemarkPYFull"`
	RemarkShortPingyin string `gorm:"column:m_nsRemarkPYShort"`
	CertificationFlag  int64  `gorm:"column:m_uiCertificationFlag"`
	Sex                int64  `gorm:"column:m_uiSex"`
	Type               int64  `gorm:"column:m_uiType"`
	ImgStatus          string `gorm:"column:m_nsImgStatus"`
	ImgKey             int64  `gorm:"column:m_uiImgKey"`
	HeadImgUrl         string `gorm:"column:m_nsHeadImgUrl"`
	HeadHDImgUrl       string `gorm:"column:m_nsHeadHDImgUrl"`
	BrandIconUrl       string `gorm:"column:m_nsBrandIconUrl"`
	AliasName          string `gorm:"column:m_nsAliasName"`
	EncodeUserName     string `gorm:"column:m_nsEncodeUserName"`
}

func (ContactPerson) TableName() string {
	return "WCContact"
}

// Message
type SQLiteSequence struct {
	Name string
	Seq  uint64
}

func (SQLiteSequence) TableName() string {
	return "sqlite_sequence"
}

type MessageContent struct {
	MesLocalID    int64             `gorm:"column:mesLocalID"`
	MesSvrID      int64             `gorm:"column:mesSvrID"`
	MsgCreateTime int64             `gorm:"column:msgCreateTime"`
	MsgContent    string            `gorm:"column:msgContent"`
	MsgStatus     int64             `gorm:"column:msgStatus"`
	MsgImgStatus  int64             `gorm:"column:msgImgStatus"`
	MessageType   model.MessageType `gorm:"column:messageType"`
	MesDes        bool              `gorm:"column:mesDes"`
	MsgSource     string            `gorm:"column:msgSource"`
	MsgVoiceText  string            `gorm:"column:msgVoiceText"`
	MsgSeq        int64             `gorm:"column:msgSeq"`
}

type HlinkMediaRecord struct {
	MediaMd5    string           `gorm:"column:mediaMd5"`
	MediaSize   int64            `gorm:"column:mediaSize"`
	InodeNumber int64            `gorm:"primary_key;column:inodeNumber"`
	ModifyTime  int64            `gorm:"column:modifyTime"`
	Detail      HlinkMediaDetail `gorm:"foreignKey:InodeNumber"`
}

func (HlinkMediaRecord) TableName() string {
	return "HlinkMediaRecord"
}

type HlinkMediaDetail struct {
	LocalId      int64  `gorm:"column:localId"`
	InodeNumber  int64  `gorm:"primary_key;column:inodeNumber"`
	RelativePath string `gorm:"column:relativePath"`
	FileName     string `gorm:"column:fileName"`
}

func (HlinkMediaDetail) TableName() string {
	return "HlinkMediaDetail"
}

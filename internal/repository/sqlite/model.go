package sqlite

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

// Message
type SQLiteSequence struct {
	Name string
	Seq  uint64
}

func (SQLiteSequence) TableName() string {
	return "sqlite_sequence"
}

type MessageContent struct {
	MesLocalID    int64  `gorm:"column:mesLocalID"`
	MesSvrID      int64  `gorm:"column:mesSvrID"`
	MsgCreateTime int64  `gorm:"column:msgCreateTime"`
	MsgContent    string `gorm:"column:msgContent"`
	MsgStatus     int64  `gorm:"column:msgStatus"`
	MsgImgStatus  int64  `gorm:"column:msgImgStatus"`
	MessageType   int64  `gorm:"column:messageType"`
	MesDes        int64  `gorm:"column:mesDes"`
	MsgSource     string `gorm:"column:msgSource"`
	MsgVoiceText  string `gorm:"column:msgVoiceText"`
	MsgSeq        int64  `gorm:"column:msgSeq"`
}

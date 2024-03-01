package repository

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

func (MessageContent) TableName(msgName string) string {
	return msgName
}

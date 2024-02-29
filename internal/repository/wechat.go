package repository

type SQLiteSequence struct {
	Name string
	Seq  uint64
}

type GroupContact struct {
	UsrName           string `gorm:"m_nsUsrName"`
	Nickname          string `gorm:"nickname"`
	Type              int64  `gorm:"m_uiType"`
	ImgStatus         string `gorm:"m_nsImgStatus"`
	HeadImgUrl        string `gorm:"m_nsHeadImgUrl"`
	ChatRoomMemList   string `gorm:"m_nsChatRoomMemList"`
	ChatRoomAdminList string `gorm:"m_nsChatRoomAdminList"`
	ChatRoomStatus    int64  `gorm:"m_uiChatRoomStatus"`
	ChatRoomVersion   int64  `gorm:"m_uiChatRoomVersion"`
	ChatRoomType      int64  `gorm:"m_uiChatRoomType"`
}

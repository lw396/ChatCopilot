package repository

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
}

func (GroupContact) TableName() string {
	return "GroupContact"
}

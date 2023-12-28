package repository

import "time"

type Article struct {
	Id         uint64
	Title      string    `json:"title"`      // 文章标题
	Content    string    `json:"content"`    // 文章内容
	Createtime time.Time `json:"createtime"` // 创建时间
	Updatetime time.Time `json:"updatetime"` // 修改时间
}

func (Article) TableName() string {
	return "pw_article"
}

package repository

import (
	"time"
)

type UserAudit struct {
	Avatar    string    `json:"avatar"`
	Nickname  string    `json:"nickname"`
	Username  string    `json:"username"`
	Sex       uint8     `json:"sex"`
	Mobile    string    `json:"mobile"`
	Work      string    `json:"work"`
	Address   string    `json:"address"`
	Birthday  time.Time `json:"birthday"`
	Photo     string    `json:"photo"`
	CardFront string    `json:"card_front"`
	CardBack  string    `json:"card_back"`
	Loginip   string    `json:"loginip"`
}

func (UserAudit) TableName() string {
	return "pw_user_audit"
}

type User struct {
	Id             uint64  `json:"id"`
	Username       string  `json:"username"`
	Nickname       string  `json:"nickname"`
	Password       string  `json:"password"`
	Salt           string  `json:"salt"`
	Mobile         string  `json:"mobile"`
	Avatar         string  `json:"avatar"`
	Maxsuccessions uint64  `json:"maxsuccessions"`
	Prevtime       uint64  `json:"prevtime"`
	Logintime      int     `json:"logintime"`
	Loginip        string  `json:"loginip"`
	Joinip         string  `json:"joinip"`
	Jointime       uint64  `json:"jointime"`
	Token          string  `json:"token"`
	Sex            uint8   `json:"sex"`
	Status         int     `json:"status"`
	Verification   string  `json:"verification"`
	IsDel          int     `json:"is_del"`
	IsType         int     `json:"is_type"`
	Top            int     `json:"top"`
	Ment           int     `json:"ment"`
	Integral       float64 `json:"integral"`
	Money          float64 `json:"money"`
	Openid         string  `json:"openid"`
	IsClassify     uint8   `json:"is_classify"`
	ClId           string  `json:"cl_id"`
	ClTitle        string  `json:"cl_title"`
	SetMealIds     string  `json:"set_meal_ids"`
	LabelId        string  `json:"label_id"`
	Label          string  `json:"label"`
	Signature      string  `json:"signature"`
	Voice          string  `json:"voice"`
	Second         uint64  `json:"second"`
	Weigh          int     `json:"weigh"`
}

func (User) TableName() string {
	return "pw_user"
}

type UserToken struct {
	Id         uint64 `json:"id"`
	UserId     uint64 `json:"user_id"`
	Token      string `json:"token"`
	ActType    uint8  `json:"act_type"`
	Createtime int64  `json:"createtime"`
	LastTime   int64  `json:"last_time"`
}

func (UserToken) TableName() string {
	return "pw_user_token"
}

type UserCollect struct {
	Id          uint64 `json:"id"`
	UserId      uint64 `json:"user_id"`
	CollClerkId uint64 `json:"coll_clerk_id"`
	Status      uint8  `json:"status"`
	CreateTime  int64  `json:"create_time"`
	UpdateTime  int64  `json:"update_time"`
}

func (UserCollect) TableName() string {
	return "pw_user_collect"
}

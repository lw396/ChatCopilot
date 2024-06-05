package service

import (
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"github.com/lw396/WeComCopilot/internal/errors"
	"github.com/lw396/WeComCopilot/internal/model"
	mysql "github.com/lw396/WeComCopilot/internal/repository/gorm"
	"github.com/lw396/WeComCopilot/internal/repository/sqlite"
	"github.com/lw396/WeComCopilot/pkg/db"
	"github.com/lw396/WeComCopilot/pkg/util"
	"gorm.io/gorm"
)

type MessageInfo struct {
	UserName string `json:"user_name"`
	Seq      uint64 `json:"seq"`
	DBName   string `json:"db_name"`
}

func (a *Service) ScanMessage(ctx context.Context, userName string) (result *MessageInfo, err error) {
	var dbName string
	var seq *sqlite.SQLiteSequence
	var name string = "Chat_" + hex.EncodeToString(util.Md5([]byte(userName)))
	for i := 0; i < 10; i++ {
		dbName = fmt.Sprintf(sqlite.MessageDB, i)
		var tx *gorm.DB
		if tx, err = a.sqlite.OpenDB(ctx, dbName); err != nil {
			return
		}
		if seq, err = a.sqlite.CheckMessageExistDB(ctx, tx, name); err != nil {
			if !db.IsRecordNotFound(err) {
				return
			}
			continue
		}
		break
	}
	if seq == nil {
		err = errors.New(errors.CodeDB, "未找到该消息")
		return
	}
	result = &MessageInfo{
		DBName:   dbName,
		UserName: userName,
		Seq:      seq.Seq,
	}
	return
}

func (a *Service) GetMessageContent(ctx context.Context, usrName string, offset int) (result []*mysql.MessageContent, err error) {
	msgName := "Chat_" + hex.EncodeToString(util.Md5([]byte(usrName)))
	result, err = a.rep.GetMessageContentList(ctx, msgName, offset)
	if err != nil {
		return
	}

	return
}

func (a *Service) convertMessageContent(ctx context.Context, msg []*sqlite.MessageContent, isGroup bool) (result []*mysql.MessageContent, err error) {
	result = make([]*mysql.MessageContent, 0)
	for _, v := range msg {
		var content string
		if content, err = a.GetHinkMedia(ctx, v, isGroup); err != nil {
			return
		}

		result = append(result, &mysql.MessageContent{
			LocalID:     v.MesLocalID,
			SvrID:       v.MesSvrID,
			CreateTime:  v.MsgCreateTime,
			Content:     content,
			Status:      v.MsgStatus,
			ImgStatus:   v.MsgImgStatus,
			MessageType: v.MessageType,
			Des:         v.MesDes,
			Source:      v.MsgSource,
			VoiceText:   v.MsgVoiceText,
			Seq:         v.MsgSeq,
		})
	}
	return
}

func (a *Service) GetHinkMedia(ctx context.Context, data *sqlite.MessageContent, isGroup bool) (result string, err error) {
	switch data.MessageType {
	case model.MsgTypeImage:
		result, err = a.HandleImage(ctx, data, isGroup)
		if err != nil {
			return
		}
	case model.MsgTypeEmoticon:
		result, err = a.HandleSticker(ctx, data, isGroup)
		if err != nil {
			return
		}

	// case model.MsgTypeVideo:

	// case model.MsgTypeVoice:

	// case model.MsgTypeMicroVideo:

	default:
		result = data.MsgContent
	}
	return
}

func (a *Service) GetMessageImage(ctx context.Context, path string) (result string, err error) {
	result = fmt.Sprintf("%s/Message/MessageTemp/%s", a.path, path)
	if _, err = os.Stat(result); err != nil {
		return
	}
	return
}

// 保存表情包路径
const StickerDir = "./data/sticker/"

func (a *Service) GetMessageSticker(ctx context.Context, path, url string) (result string, err error) {
	result = StickerDir + path
	if _, err = os.Stat(result); err != nil && os.IsExist(err) {
		return
	}
	if os.IsNotExist(err) {
		if err = a.CacheSticker(ctx, result, url); err != nil {
			return
		}
	}

	return
}

func (a *Service) CacheSticker(ctx context.Context, path, url string) (err error) {
	url = strings.ReplaceAll(url, "\\u0026", "&")
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if _, err = os.Stat(StickerDir); err != nil && os.IsExist(err) {
		return
	}
	if os.IsNotExist(err) {
		if err = os.MkdirAll(StickerDir, fs.ModePerm); err != nil {
			return
		}
	}
	if err = os.WriteFile(path, content, 0644); err != nil {
		return
	}
	return
}

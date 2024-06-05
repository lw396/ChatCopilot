package service

import (
	"context"
	"encoding/xml"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/lw396/WeComCopilot/internal/model"
	"github.com/lw396/WeComCopilot/internal/repository/sqlite"
	"github.com/lw396/WeComCopilot/pkg/db"
	"howett.net/plist"
)

type MediaMessage struct {
	Sender string `json:"sender"`
	Path   string `json:"path"`
	Url    string `json:"url"`
	Md5    string `json:"md5"`
}

type ImageMessageData struct {
	XMLName xml.Name `xml:"msg"`
	Img     struct {
		Md5 string `xml:"md5,attr"`
	} `xml:"img"`
}

type VideoMessageData struct {
	XMLName xml.Name `xml:"msg"`
	Video   struct {
		Md5 string `xml:"md5,attr"`
	} `xml:"videomsg"`
}

func (a *Service) HandleImage(ctx context.Context, message *sqlite.MessageContent, isGroup bool) (result *MediaMessage, err error) {
	var data ImageMessageData
	if err = xml.Unmarshal([]byte(message.MsgContent), &data); err != nil {
		return
	}
	var sender, path string
	if message.MesDes && isGroup {
		sender = strings.Split(message.MsgContent, ":")[0]
	}

	if data.Img.Md5 != "" {
		if err = a.ConnectDB(ctx, sqlite.HlinkDB); err != nil {
			return
		}

		hlink := &sqlite.HlinkMediaRecord{}
		hlink, err = a.sqlite.GetHinkMediaByMediaMd5(ctx, data.Img.Md5)
		if err != nil && !db.IsRecordNotFound(err) {
			return
		}

		if !db.IsRecordNotFound(err) {
			path = hlink.Detail.RelativePath + hlink.Detail.FileName
		}
	}

	result = &MediaMessage{
		Sender: sender,
		Path:   path,
		Md5:    data.Img.Md5,
	}
	return
}

type StickerMessageData struct {
	XMLName xml.Name `xml:"msg"`
	Sticker struct {
		Md5 string `xml:"md5,attr"`
		Url string `xml:"cdnurl,attr"`
	} `xml:"emoji"`
}

func (a *Service) HandleSticker(ctx context.Context, message *sqlite.MessageContent, isGroup bool) (result *MediaMessage, err error) {
	var data StickerMessageData
	if err = xml.Unmarshal([]byte(message.MsgContent), &data); err != nil {
		return
	}

	var sender string
	if message.MesDes && isGroup {
		sender = strings.Split(message.MsgContent, ":")[0]
	}

	var url string
	if data.Sticker.Md5 != "" {
		if data.Sticker.Url != "" {
			url = strings.ReplaceAll(data.Sticker.Url, "amp;", "")
		}
		if data.Sticker.Url == "" {
			if url, err = a.GetStickerFavArchive(ctx, data.Sticker.Md5); err != nil {
				return
			}
		}
	}

	result = &MediaMessage{
		Sender: sender,
		Path:   data.Sticker.Md5,
		Md5:    data.Sticker.Md5,
		Url:    url,
	}
	return
}

// 获取收藏表情包
func (a *Service) GetStickerFavArchive(ctx context.Context, md5 string) (result string, err error) {
	f, err := os.Open(a.path + "/Stickers/fav.archive")
	if err != nil {
		return
	}
	defer f.Close()

	var data map[string]any
	err = plist.NewDecoder(f).Decode(&data)
	if err != nil {
		return
	}
	var _url *url.URL
	for _, item := range data["$objects"].([]any) {
		str, succ := item.(string)
		if !succ {
			continue
		}
		if !strings.Contains(str, md5) {
			continue
		}
		_url, err = url.ParseRequestURI(str)
		if err != nil {
			continue
		}
		break
	}

	if _url == nil {
		return
	}
	result = _url.String()
	return
}

func (a *Service) HandleVideo(ctx context.Context, message *sqlite.MessageContent, isGroup bool) (result string, err error) {
	return
}

type RecordUndownloadedFileParams struct {
	Md5         string
	Sender      string
	LocalID     int64
	MessageType model.MessageType
	CreatedAt   time.Time
}

func (a *Service) recordUndownloadedFile(ctx context.Context, params []RecordUndownloadedFileParams) (err error) {
	_params := make([]RecordUndownloadedFileParams, 0)
	if _, err = a.redis.Get(ctx, SyncTaskUnloadedFile, &_params); err != nil {
		return
	}

	var now = time.Now()
	for _, param := range params {
		if !param.CreatedAt.After(now) {
			continue
		}
		params = append(params, param)
	}
	if err = a.redis.Set(ctx, SyncTaskUnloadedFile, params, time.Minute*10); err != nil {
		return
	}
	return
}

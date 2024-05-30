package service

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"strings"

	"github.com/lw396/WeComCopilot/internal/model"
	"github.com/lw396/WeComCopilot/internal/repository/sqlite"
)

type MediaMessage struct {
	Sender      string            `json:"sender"`
	Path        string            `json:"path"`
	Url         string            `json:"url"`
	MessageType model.MessageType `json:"message_type"`
}

type ImageMessageData struct {
	XMLName xml.Name `xml:"msg"`
	Img     struct {
		Md5 string `xml:"md5,attr"`
	} `xml:"img"`
}

func (a *Service) HandleImage(ctx context.Context, message *sqlite.MessageContent, isGroup bool) (result string, err error) {
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
		var hlink *sqlite.HlinkMediaRecord
		if hlink, err = a.sqlite.GetHinkMediaByMediaMd5(ctx, data.Img.Md5); err != nil {
			return
		}
		path = hlink.Detail.RelativePath + hlink.Detail.FileName
	}

	_result, err := json.Marshal(&MediaMessage{
		Sender:      sender,
		Path:        path,
		MessageType: model.MsgTypeImage,
	})
	if err != nil {
		return
	}

	result = string(_result)
	return
}

type StickerMessageData struct {
	XMLName xml.Name `xml:"msg"`
	Sticker struct {
		Md5 string `xml:"md5,attr"`
		Url string `xml:"cdnurl,attr"`
	} `xml:"emoji"`
}

func (a *Service) HandleSticker(ctx context.Context, message *sqlite.MessageContent, isGroup bool) (result string, err error) {
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
			if url, err = a.sqlite.GetStickerFavArchive(ctx, data.Sticker.Md5); err != nil {
				return
			}
		}
	}
	_result, err := json.Marshal(&MediaMessage{
		Sender:      sender,
		Path:        data.Sticker.Md5,
		Url:         url,
		MessageType: model.MsgTypeEmoticon,
	})
	if err != nil {
		return
	}

	result = string(_result)
	return
}

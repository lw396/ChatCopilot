package service

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"strings"

	"github.com/lw396/WeComCopilot/internal/model"
	"github.com/lw396/WeComCopilot/internal/repository/sqlite"
)

// ImageMessageData 图片消息结构体
type ImageMessageData struct {
	XMLName xml.Name `xml:"msg"`
	Img     struct {
		Text           string `xml:",chardata"`
		AesKey         string `xml:"aeskey,attr"`
		EnCryVer       string `xml:"encryver,attr"`
		CdnThumbAesKey string `xml:"cdnthumbaeskey,attr"`
		CdnThumbUrl    string `xml:"cdnthumburl,attr"`
		CdnThumbLength string `xml:"cdnthumblength,attr"`
		CdnThumbHeight string `xml:"cdnthumbheight,attr"`
		CdnThumbWidth  string `xml:"cdnthumbwidth,attr"`
		CdnMidHeight   string `xml:"cdnmidheight,attr"`
		CdnMidWidth    string `xml:"cdnmidwidth,attr"`
		CdnHdHeight    string `xml:"cdnhdheight,attr"`
		CdnHdWidth     string `xml:"cdnhdwidth,attr"`
		CdnMidImgUrl   string `xml:"cdnmidimgurl,attr"`
		Length         int64  `xml:"length,attr"`
		CdnBigImgUrl   string `xml:"cdnbigimgurl,attr"`
		HdLength       int64  `xml:"hdlength,attr"`
		Md5            string `xml:"md5,attr"`
	} `xml:"img"`
}

type MediaMessage struct {
	Sender      string            `json:"sender"`
	Path        string            `json:"path"`
	MessageType model.MessageType `json:"message_type"`
}

func (a *Service) HandleImage(ctx context.Context, content string, isDes, isGroup bool) (result string, err error) {
	var data ImageMessageData
	if err = xml.Unmarshal([]byte(content), &data); err != nil {
		return
	}
	var sender, path string
	if isDes && isGroup {
		sender = strings.Split(content, ":")[0]
	}

	if data.Img.Md5 != "" {
		if err = a.ConnectDB(ctx, sqlite.HlinkDB); err != nil {
			return
		}
		var hlink *sqlite.HlinkMediaRecord
		hlink, err = a.sqlite.GetHinkMediaByMediaMd5(ctx, data.Img.Md5)
		if err != nil {
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

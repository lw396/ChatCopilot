package sqlite

import (
	"context"
	"net/url"
	"os"
	"strings"

	"howett.net/plist"
)

func (s *SQLite) GetStickerFavArchive(ctx context.Context, md5 string) (result string, err error) {
	f, err := os.Open(s.path + "/" + FavArchive)
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

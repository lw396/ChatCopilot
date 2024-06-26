package audio

import "C"
import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const ResourceDir = "./data"

func Decoder(folder, fileName string) (err error) {
	var (
		file   = strings.TrimSuffix(fileName, filepath.Ext(fileName))
		input  = filepath.Join(folder, fileName)
		output = filepath.Join(ResourceDir, file+".pcm")
	)

	dir := filepath.Dir(output)
	if _, err = os.Stat(dir); err != nil && !os.IsNotExist(err) {
		return
	}
	if os.IsNotExist(err) {
		if err = os.MkdirAll(dir, fs.ModePerm); err != nil {
			return
		}
	}
	if err = SilkToPcm(input, output); err != nil {
		return
	}

	input = output
	output = filepath.Join(ResourceDir, file+".wav")
	if err = PcmToWav(input, output); err != nil {
		return
	}

	if err = os.Remove(input); err != nil {
		return
	}
	return
}

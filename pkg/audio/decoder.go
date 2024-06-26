package silkv3decoder

/*
#cgo CFLAGS: -I ./silk/src
#cgo CFLAGS: -I ./silk/interface
#cgo CFLAGS: -I ./silk/test
#cgo LDFLAGS: -lm
#cgo LDFLAGS: ./silk/libSKP_SILK_SDK.a

#include <stdlib.h>
#include <string.h>
#include <Decoder.c>
*/
import "C"
import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"unsafe"

	"github.com/lw396/WeComCopilot/internal/errors"
)

const ResourceDir = "./data"

func SilkToPcm(folder, fileName string) (err error) {
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

	fmt.Println(input, output)
	args := []string{"", input, output}
	argc := C.int(len(args))
	argv := make([]*C.char, len(args))
	for i, arg := range args {
		argv[i] = C.CString(arg)
		defer C.free(unsafe.Pointer(argv[i]))
	}

	if errcode := C.decoder(argc, (**C.char)(unsafe.Pointer(&argv[0]))); errcode != 0 {
		err = errors.New(errors.CodeGeneral, "Silk to pem error")
		return
	}
	return
}

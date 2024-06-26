package audio

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
	"unsafe"

	"github.com/lw396/WeComCopilot/internal/errors"
)

func SilkToPcm(input, output string) (err error) {
	var (
		args = []string{"", input, output}
		argc = C.int(len(args))
		argv = make([]*C.char, len(args))
	)
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

package audio

/*
#cgo CFLAGS: -I ./silk/src
#cgo CFLAGS: -I ./silk/interface
#cgo CFLAGS: -I ./silk/test
#cgo LDFLAGS: -lm
#cgo LDFLAGS: pkg/audio/silk/libSKP_SILK_SDK.a

#include <stdlib.h>
#include <string.h>
#include <Decoder.c>
*/
import "C"
import (
	"github.com/lw396/WeComCopilot/internal/errors"
)

func SilkToPcm(input, output string) (err error) {
	if errcode := C.decoder(C.CString(input), C.CString(output)); errcode != 0 {
		err = errors.New(errors.CodeGeneral, "Silk to pem error")
		return
	}
	return
}

package silkv3decoder

/*
#cgo CFLAGS: -I ./silk/interface
#cgo CFLAGS: -I ./silk/src
#cgo LDFLAGS: -lm
#cgo LDFLAGS: ./silk/libSKP_SILK_SDK.a

#include <stdlib.h>
#include <string.h>
#include "SKP_Silk_SDK_API.h"
#include "SKP_Silk_control.h"

#include "SKP_Silk_SigProc_FIX.h"
*/
import "C"
import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	MAX_LBRR_DELAY = 2
)

func Decoder(inputFolder, inputFile, outputFolder string) (err error) {
	var (
		nBytes          C.SKP_int16
		decSizeBytes    C.SKP_int32
		DecControl      C.SKP_SILK_SDK_DecControlStruct
		nBytesPerPacket [MAX_LBRR_DELAY + 1]C.SKP_int16
		payloadEnd      *C.SKP_int8 = nil
		totPackets      C.SKP_int32
	)
	file, err := os.Open(inputFolder + "/" + inputFile)
	if err != nil {
		return
	}
	defer file.Close()

	if err = CheckSilkHeader(file); err != nil {
		return
	}

	newFile, err := CreateOutputFile(inputFile, outputFolder)
	if err != nil {
		return
	}

	DecControl = C.SKP_SILK_SDK_DecControlStruct{
		API_sampleRate:  24000,
		framesPerPacket: 1,
	}

	// var decSizeBytes C.SKP_int32
	// ret := C.SKP_Silk_SDK_Get_Decoder_Size(&decSizeBytes)
	// fmt.Println("size: ", ret)

	// psDec := make([]byte, decSizeBytes)

	// ret = C.SKP_Silk_SDK_InitDecoder(unsafe.Pointer(&psDec))
	// fmt.Println("size: ", ret)
	for i := 0; i < MAX_LBRR_DELAY; i++ {
		err = binary.Read(file, binary.LittleEndian, &nBytes)
		if err != nil {
			break
		}

		payload := make([]byte, nBytes)
		counter, err := file.Read(payload)
		if err != nil || int16(counter) < int16(nBytes) {
			break
		}

		nBytesPerPacket[i] = nBytes
		payloadEnd = append(payloadEnd, payload...)
		totPackets++
	}

	return
}

func CheckSilkHeader(file *os.File) (err error) {
	header := make([]byte, 50)
	if _, err = io.ReadFull(file, header[:1]); err != nil {
		return
	}

	silk := "#!SILK_V3"
	if string(header[:1]) != "\x02" {
		silk = "!SILK_V3"
	}
	if _, err = io.ReadFull(file, header[:len(silk)]); err != nil {
		return
	}
	header[len(silk)] = '\x00'
	if string(header[:len(silk)]) != silk {
		err = errors.New("not a silk file")
		return
	}

	return
}

func CreateOutputFile(inputFile, outputFolder string) (result *os.File, err error) {
	_fileName := strings.SplitAfter(inputFile, ".")
	fileName := strings.Join(_fileName[:len(_fileName)-1], "")
	fullFile := fmt.Sprintf("%s/%s%s", outputFolder, fileName, "pcm")

	result, err = os.Create(fullFile)
	if err != nil {
		return
	}
	defer result.Close()

	return
}

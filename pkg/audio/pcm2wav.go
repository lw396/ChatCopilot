package audio

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"
)

type Head struct {
	ChunkID   [4]byte // 内容为"RIFF"
	ChunkSize uint32  // wav文件的字节数, 不包含ChunkID和ChunkSize这8个字节）
	Format    [4]byte // 内容为WAVE
}

type Fmt struct {
	Subchunk1ID   [4]byte // 内容为"fmt "
	Subchunk1Size uint32  // Fmt所占字节数，为16
	AudioFormat   uint16  // 存储音频的编码格式，pcm为1
	NumChannels   uint16  // 通道数, 单通道为1,双通道为2
	SampleRate    uint32  // 采样率，如8k, 44.1k等
	ByteRate      uint32  // 每秒存储的byte数，其值=SampleRate * NumChannels * BitsPerSample/8
	BlockAlign    uint16  // 块对齐大小，其值=NumChannels * BitsPerSample/8
	BitsPerSample uint16  // 每个采样点的bit数，一般为8,16,32等。
}

type Data struct {
	Subchunk2ID   [4]byte // 内容为"data"
	Subchunk2Size uint32  // 内容为接下来的正式的数据部分的字节数，其值=NumSamples * NumChannels * BitsPerSample/8
}

type WavHead struct {
	Head
	Fmt
	Data
}

func New(numChannels uint16, sampleRate uint32, bitsPerSample uint16, wavLen uint32) *WavHead {
	head := &WavHead{
		Head: Head{
			ChunkSize: uint32(36 + wavLen),
		},
		Fmt: Fmt{
			Subchunk1Size: 16,
			AudioFormat:   1,
			NumChannels:   numChannels,
			SampleRate:    sampleRate,
			ByteRate:      sampleRate * uint32(numChannels) * uint32(bitsPerSample) / 8,
			BlockAlign:    numChannels * bitsPerSample / 8,
			BitsPerSample: bitsPerSample,
		},
		Data: Data{
			Subchunk2Size: uint32(wavLen),
		},
	}

	copy(head.ChunkID[:], "RIFF")
	copy(head.Format[:], "WAVE")
	copy(head.Subchunk1ID[:], "fmt ")
	copy(head.Subchunk2ID[:], "data")
	return head
}

func (wh *WavHead) Marshal() ([]byte, error) {
	buf := &bytes.Buffer{}
	if err := binary.Write(buf, binary.LittleEndian, *wh); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func PcmToWav(file, wavFile string) (err error) {
	pcmFile, err := os.Open(file)
	if err != nil {
		return
	}
	defer pcmFile.Close()

	pcmFi, err := pcmFile.Stat()
	if err != nil {
		return
	}

	wavFd, err := os.Create(wavFile)
	if err != nil {
		return
	}
	defer wavFd.Close()

	wavHead := New(uint16(1), uint32(24000), uint16(16), uint32(pcmFi.Size()))
	head, err := wavHead.Marshal()
	if err != nil {
		return
	}

	wavFd.Write(head)
	io.Copy(wavFd, pcmFile)

	return
}

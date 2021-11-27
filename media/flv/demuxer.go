package flv

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

type Demuxer struct {
	rd          io.Reader
	lastTagSize uint32
}

type MediaWriter interface {
func Write
}

var ErrInvalidTag = errors.New("flv: invalid tag")
var ErrInvalidHeader = errors.New("flv: invalid header")

func NewDecoder(r io.Reader) *Demuxer {
	return &Demuxer{rd: r}
}


// ParseHeader 解析直到
func (fd *Demuxer) ParseHeader() error {
	const (
		headerConst = "FLV\x01\x05\x00\x00\x00\x09"
	)

	var header [len(headerConst)]byte
	_, err := io.ReadFull(fd.rd, header[:])
	if err != nil {
		return err
	}
	if bytes.Equal(header[:], []byte(headerConst)) {
		return nil
	}

	if header[0] != 'F' || header[1] != 'L' || header[2] != 'V' {
		return ErrInvalidHeader
	}

	if header[3] != 1 {
		return fmt.Errorf("flv: 不支持的版本：%d", header[3])
	}

	if header[4] != 0b00000101 {
		return fmt.Errorf("flv: 仅支持音+视频文件")
	}

	return fmt.Errorf("flv: 解析头部时出现未知错误：%x", header[:9])
}

func (fd *Demuxer) NextChunk() (*Chunk, error {
	var tagType uint8
	var tagSize uint32
	var err error

	if fd.lastTagSize > 0 {
		tagSize = fd.lastTagSize
		fd.lastTagSize = 0
	} else {
		err = fd.readTagSize(&tagSize)
		if err != nil {
			return nil, err
		}
	}

	err = fd.readTagType(&tagType)
	if err != nil {
		return nil, err
	}

	chunk := &Chunk{
		Type: tagType,
		Data: make([]byte, tagSize),
	}

	_, err = io.ReadFull(fd.rd, chunk.Data)
	if err != nil {
		return nil, err
	}

	return chunk, nil
}

package mp4

import "os"

const (
	ftyp = "\u0000\u0020ftypisom\u0200isomiso2avc1mp41"
)

type Muxer struct {
	moov atom
	mdat atom
}

type Info struct {
	Width, Height int
}

func NewMP4(info Info) {
	os.File{}.WriteAt()
}

package media

type Writer interface {
	SetMediaInfo()
	WriteFrame(id uint8, timestamp uint64, data []byte)
}

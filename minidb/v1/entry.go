package v1

const entryHeaderSize = 10

const (
	PUT uint16 = iota
	DEL
)

// Entry 写入文件的记录
type Entry struct {
	Key       []byte
	Value     []byte
	KeySize   uint32
	ValueSize uint32
	Mark      uint16
}

func (e *Entry) GetSize() int64 {
	return int64(entryHeaderSize + e.KeySize + e.ValueSize)
}

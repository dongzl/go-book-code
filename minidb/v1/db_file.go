package v1

import "os"

const (
	FileName = "minidb.data"
)

// DBFile 数据文件定义
type DBFile struct {
	File   *os.File
	Offset int64
}

func (df *DBFile) Read(offset int64) (*Entry, error) {
	return nil, nil
}

// NewDBFile 创建一个新的数据文件
func NewDBFile(path string) (*DBFile, error) {
	fileName := path + string(os.PathSeparator) + FileName
	return newInternal(fileName)
}

func newInternal(fileName string) (*DBFile, error) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	stat, err := os.Stat(fileName)
	if err != nil {
		return nil, err
	}
	return &DBFile{File: file, Offset: stat.Size()}, nil
}

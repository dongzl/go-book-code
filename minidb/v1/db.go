package v1

import (
	"io"
	"os"
	"sync"
)

type MiniDB struct {
	indexes map[string]int64 // 内存中的索引信息
	dbFile  *DBFile          // 数据文件
	dirPath string           // 数据目录
	mu      sync.RWMutex
}

// Open 开启一个数据库实例
func Open(dirPath string) (*MiniDB, error) {
	// 如果数据库目录不存在，则新建一个
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return nil, err
		}
	}

	// 加载数据文件
	dbFile, err := NewDBFile(dirPath);
	if err != nil {
		return nil, err
	}
	
	db := &MiniDB{
		dbFile: dbFile,
		indexes: make(map[string]int64),
		dirPath: dirPath,
	}

	// 加载索引
	db.loadIndexesFromFile()
	return db, nil
}

func (db *MiniDB) loadIndexesFromFile() {
	if db.dbFile == nil {
		return
	}
	var offset int64
	for {
		e, err := db.dbFile.Read(offset)
		if err != nil {
			// 读取完毕
			if err == io.EOF {
				break
			}
			return
		}

		// 设置索引状态
		db.indexes[string(e.Key)] = offset
		if e.Mark == DEL {
			// 删除内存中的 key
			delete(db.indexes, string(e.Key))
		}

		offset += e.GetSize()
	}
	return
}

package pool

type DBConn struct {
	idleTime int
}

// 新建数据库连接
func NewDBConn() *DBConn {
	return &DBConn{idleTime: 0}
}

// 关闭数据库连接
func (d *DBConn) Close() {}

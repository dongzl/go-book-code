package main

import "net"

func main()  {
	// 监听本地 9090 端口
	socket, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		panic(err)
	}

	defer socket.Close()

	for {
		// 处理连接请求
		conn, err := socket.Accept();
		if err != nil {
			panic(err)
		}
		// 处理已经成功建立连接的请求
		go handleRequest(conn)
	}
}

// 处理已经成功建立连接的请求
func handleRequest(conn net.Conn)  {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		// 读取请求数据
		size, err := conn.Read(buf)
		if err != nil {
			return
		}
		// 回写相应数据
		conn.Write(buf[:size])
	}
}

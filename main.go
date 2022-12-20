package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

// rain-service http 协议文件下载服务
// 基于 json-rpc 1.0 协议

var ()

func main() {
	var err error

	// 初始化数据库
	sqliteCtl, err := NewSqliteControl()
	if err != nil {
		panic(err)
	}

	// 读取数据
	rowTables, err := sqliteCtl.GetRows()
	if err != nil {
		panic(err)
	}
	rows := make([]*Row, 0, len(rowTables))
	for _, v := range rowTables {
		rows = append(rows, v.Row())
	}

	// 注册 rpc
	service := NewService(rows)
	rpc.Register(service)

	// 数据库自动存储
	go func() {
		for {
			time.Sleep(time.Second * 10)
			sqliteCtl.Update(service.Rows)
		}
	}()

	// 程序退出监听

	http.HandleFunc("/json", JsonRPC)

	err = http.ListenAndServe(":7362", nil)
	if err != nil {
		panic(err)
	}
}

func JsonRPC(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// 声明一个客户端连接对象
	var conn io.ReadWriteCloser = struct {
		io.Writer
		io.ReadCloser
	}{
		ReadCloser: req.Body,
		Writer:     w,
	}
	rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
}

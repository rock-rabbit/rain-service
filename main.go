package main

import (
	"context"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// rain-service http 协议文件下载服务
// 基于 json-rpc 1.0 协议

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

	// 启动 http 服务
	http.HandleFunc("/json", JsonRPC)
	server := &http.Server{Addr: ":7362", Handler: nil}
	go server.ListenAndServe()

	// 程序退出监听
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 退出服务
	server.Shutdown(context.Background())
	// 保存数据
	sqliteCtl.Update(service.Rows)

	os.Exit(0)
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

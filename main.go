package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// rain-service http 协议文件下载服务
// 基于 json-rpc 1.0 协议

func main() {
	var err error

	rpc.Register(NewService())

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

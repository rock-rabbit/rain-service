## rain-service -json-rpc 文件下载服务
rain-service 是基于 rain 的 http 协议文件下载服务

## 特色
* rpc-json 1.0
* 持久化存储

## 方法
``` golang
// 新增下载
Service.AddUri

// 下载列表
Service.GetRows
```
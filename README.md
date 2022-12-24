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

// 开始下载
Service.Start

// 暂停下载
Service.Pause

// 删除
Service.Delete
```

## 下载状态
``` golang
const (
	// 运行中
	STATUS_RUNNING Status = "running"
	// 暂停
	STATUS_PAUSE Status = "pause"
	// 完成
	STATUS_FINISH Status = "finish"
	// 错误
	STATUS_ERROR Status = "error"
	// 回收站
	STATUS_COVERY Status = "covery"
)
```

## 方法介绍

Service.AddUri - 新增下载

``` golang
// 发送
{
    "method": "Service.AddUri",
    "params": [
        {
            "URI": "https://images.unsplash.com/photo-1531554694128-c4c6665f59c2?ixlib=rb-4.0.3&dl=zhang-kaiyv-dhiOkqjewAM-unsplash.jpg&q=80&fm=jpg&crop=entropy&cs=tinysrgb",
            "Outdir": "./tmp",
            "Outname": "test.jpg"
        }
    ],
    "id": 0
}

// 返回
{
    "id": 0,
    "result": {
        "UUID": "affe8537-58e6-4626-8092-117e1a9b9199",
        "Status": "running",
        "URI": "https://images.unsplash.com/photo-1531554694128-c4c6665f59c2?ixlib=rb-4.0.3&dl=zhang-kaiyv-dhiOkqjewAM-unsplash.jpg&q=80&fm=jpg&crop=entropy&cs=tinysrgb",
        "Outdir": "./tmp",
        "Outname": "test2.jpg",
        "Header": null,
        "Stat": null,
        "Error": "",
        "Progress": 0,
        "CreateTime": "2022-12-21T14:12:18.761756+08:00",
        "UpdateTime": "2022-12-21T14:12:18.761757+08:00",
        "SyncTime": "0001-01-01T00:00:00Z"
    },
    "error": null
}
```

Service.GetRows - 下载列表

``` golang
// 发送
{
    "method": "Service.GetRows",
    "params": [
        {
            "Status":""
        }
    ],
    "id": 0
}

// 返回
{
    "id": 0,
    "result": [
        {
            "UUID": "70b975d7-f82e-417d-906f-3ef6d444db0f",
            "Status": "finish",
            "URI": "https://sample-videos.com/img/Sample-jpg-image-10mb.jpg",
            "Outdir": "./tmp",
            "Outname": "test.jpg",
            "Header": null,
            "Stat": {
                "Status": 2,
                "TotalLength": 10506316,
                "CompletedLength": 10506316,
                "DownloadSpeed": 4395084,
                "EstimatedTime": 0,
                "Progress": 100,
                "Outpath": "/Users/rockrabbit/projects/rain-service/tmp/test.jpg",
                "Error": null
            },
            "Error": "",
            "Progress": 100,
            "CreateTime": "2022-12-20T23:30:49.808749+08:00",
            "UpdateTime": "2022-12-21T14:01:47.219222+08:00",
            "SyncTime": "2022-12-21T14:01:51.700697+08:00"
        },
        {
            "UUID": "affe8537-58e6-4626-8092-117e1a9b9199",
            "Status": "running",
            "URI": "https://images.unsplash.com/photo-1531554694128-c4c6665f59c2?ixlib=rb-4.0.3&dl=zhang-kaiyv-dhiOkqjewAM-unsplash.jpg&q=80&fm=jpg&crop=entropy&cs=tinysrgb",
            "Outdir": "./tmp",
            "Outname": "test2.jpg",
            "Header": null,
            "Stat": {
                "Status": 1,
                "TotalLength": 4171025,
                "CompletedLength": 664128,
                "DownloadSpeed": 0,
                "EstimatedTime": 214000000000,
                "Progress": 15,
                "Outpath": "/Users/rockrabbit/projects/rain-service/tmp/test2.jpg",
                "Error": null
            },
            "Error": "",
            "Progress": 15,
            "CreateTime": "2022-12-21T14:12:18.761756+08:00",
            "UpdateTime": "2022-12-21T14:13:18.218769+08:00",
            "SyncTime": "2022-12-21T14:13:21.791091+08:00"
        }
    ],
    "error": null
}
```

Service.Start - 开始下载


``` golang
// 发送
{
    "method": "Service.Start",
    "params": [
       "affe8537-58e6-4626-8092-117e1a9b9199"
    ],
    "id": 0
}

// 返回
{
    "id": 0,
    "result": {
        "UUID": "affe8537-58e6-4626-8092-117e1a9b9199",
        "Status": "running",
        "URI": "https://images.unsplash.com/photo-1531554694128-c4c6665f59c2?ixlib=rb-4.0.3&dl=zhang-kaiyv-dhiOkqjewAM-unsplash.jpg&q=80&fm=jpg&crop=entropy&cs=tinysrgb",
        "Outdir": "./tmp",
        "Outname": "test2.jpg",
        "Header": null,
        "Stat": {
            "Status": 2,
            "TotalLength": 4171025,
            "CompletedLength": 664128,
            "DownloadSpeed": 0,
            "EstimatedTime": 214000000000,
            "Progress": 15,
            "Outpath": "/Users/rockrabbit/projects/rain-service/tmp/test2.jpg",
            "Error": null
        },
        "Error": "",
        "Progress": 15,
        "CreateTime": "2022-12-21T14:12:18.761756+08:00",
        "UpdateTime": "2022-12-21T14:18:36.390109+08:00",
        "SyncTime": "2022-12-21T14:17:51.827074+08:00"
    },
    "error": null
}
```

Service.Pause - 暂停下载

``` golang
// 发送
{
    "method": "Service.Pause",
    "params": [
       "069e982c-9ebe-47c3-8c10-f878c3854292"
    ],
    "id": 0
}

// 返回
{
    "id": 0,
    "result": {
        "UUID": "069e982c-9ebe-47c3-8c10-f878c3854292",
        "Status": "pause",
        "URI": "https://images.unsplash.com/photo-1531554694128-c4c6665f59c2?ixlib=rb-4.0.3&dl=zhang-kaiyv-dhiOkqjewAM-unsplash.jpg&q=80&fm=jpg&crop=entropy&cs=tinysrgb",
        "Outdir": "./tmp",
        "Outname": "test3.jpg",
        "Header": null,
        "Stat": {
            "Status": 1,
            "TotalLength": 4171025,
            "CompletedLength": 163840,
            "DownloadSpeed": 16384,
            "EstimatedTime": 244000000000,
            "Progress": 3,
            "Outpath": "/Users/rockrabbit/projects/rain-service/tmp/test3.jpg",
            "Error": null
        },
        "Error": "",
        "Progress": 3,
        "CreateTime": "2022-12-21T14:20:03.630539+08:00",
        "UpdateTime": "2022-12-21T14:20:12.367825+08:00",
        "SyncTime": "2022-12-21T14:20:11.855748+08:00"
    },
    "error": null
}
```
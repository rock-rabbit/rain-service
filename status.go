package main

// Status 状态
type Status string

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

// Is 对比是否相同
func (s Status) Is(v Status) bool {
	return s == v
}

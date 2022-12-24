package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/rock-rabbit/rain"
)

// Row 下载任务
type Row struct {
	UUID   string
	Status Status

	URI     string
	Outdir  string
	Outname string
	Header  map[string]string

	Stat *rain.Stat

	Error    string
	Progress int

	CreateTime time.Time
	UpdateTime time.Time

	SyncTime time.Time

	rainCtl *rain.RainControl
}

// RowEvent 下载事件监听
type RowEvent struct {
	*Row
}

// Change 下载进度事件
func (r *RowEvent) Change(stat *rain.Stat) {
	r.Stat = stat

	if r.Progress != stat.Progress {
		r.Progress = stat.Progress
		r.UpdateTime = time.Now()
	}

	if stat.Status == rain.STATUS_FINISH {
		if stat.Error != nil {
			r.SetError(stat.Error)
		} else if stat.Progress == 100 {
			r.SetStatus(STATUS_FINISH)
		}
	}
}

// Table 转为表数据
func (r *Row) Table() *RowTable {
	headerJson, _ := json.Marshal(r.Header)
	status := r.Status
	if status == STATUS_RUNNING {
		status = STATUS_PAUSE
	}
	return &RowTable{
		UUID:       r.UUID,
		Status:     string(status),
		URI:        r.URI,
		Outdir:     r.Outdir,
		Outname:    r.Outname,
		Header:     string(headerJson),
		Progress:   r.Progress,
		Error:      r.Error,
		CreateTime: r.CreateTime,
		UpdateTime: r.UpdateTime,
	}
}

// SetStatus 设置状态
func (r *Row) SetStatus(s Status) {
	r.Status = s
	r.UpdateTime = time.Now()
}

// SetError 设置错误信息
func (r *Row) SetError(err error) {
	if err == context.Canceled {
		return
	}

	r.SetStatus(STATUS_ERROR)
	r.Error = err.Error()
}

// Delete 删除下载
func (r *Row) Delete() error {
	if r.Status.Is(STATUS_RUNNING) {
		return errors.New("cannot delete because the status is running")
	}

	r.SetStatus(STATUS_COVERY)
	return nil
}

// Pause 暂停下载
func (r *Row) Pause() error {
	if !r.Status.Is(STATUS_RUNNING) {
		return errors.New("cannot pause because the status not is running")
	}

	r.rainCtl.Close()
	r.SetStatus(STATUS_PAUSE)

	return nil
}

// Start 开始下载资源
func (r *Row) Start() error {

	// 检查是否可以开始下载
	if r.Status.Is(STATUS_RUNNING) {
		return errors.New("cannot start because the status is running")
	}

	r.SetStatus(STATUS_RUNNING)

	opts := []rain.OptionFunc{
		rain.WithEvent(&RowEvent{r}),
	}

	if r.Outdir != "" {
		opts = append(opts, rain.WithOutdir(r.Outdir))
	}

	if r.Outname != "" {
		opts = append(opts, rain.WithOutname(r.Outname))
	}

	if len(r.Header) != 0 {
		opts = append(opts, rain.WithHeader(func(h http.Header) {
			for k, v := range r.Header {
				h.Set(k, v)
			}
		}))
	}

	r.rainCtl = rain.New(r.URI, opts...)
	go func() {
		err := r.rainCtl.Start()
		if err != nil {
			r.SetError(err)
		}
	}()

	return nil
}

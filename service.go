package main

import (
	"sync"
)

type Service struct {
	Rows map[string]*Row
	mux  sync.Mutex
}

const (
	// 运行中
	STATUS_RUNNING = "running"
	// 暂停
	STATUS_PAUSE = "pause"
	// 完成
	STATUS_FINISH = "finish"
	// 错误
	STATUS_ERROR = "error"
	// 回收站
	STATUS_COVERY = "covery"
)

func NewService() *Service {
	return &Service{
		Rows: make(map[string]*Row),
		mux:  sync.Mutex{},
	}
}

func (s *Service) GetRows(args *struct{ Status string }, reply *[]*Row) error {
	statusVal := args.Status
	for _, v := range s.Rows {
		if statusVal != "" && statusVal != v.Status {
			continue
		}
		*reply = append(*reply, v)
	}
	return nil
}

func (s *Service) AddUri(args *struct {
	URI     string
	Header  map[string]string
	Outdir  string
	Outname string
}, reply *bool) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	uuid := NewUUID()
	s.Rows[uuid] = &Row{
		Status:   STATUS_RUNNING,
		UUID:     uuid,
		URI:      args.URI,
		Progress: 0,
		Outdir:   args.Outdir,
		Outname:  args.Outname,
		Header:   args.Header,
	}
	s.Rows[uuid].Start()

	*reply = true
	return nil
}

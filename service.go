package main

import (
	"errors"
	"sync"
	"time"
)

// Service rpc 服务
type Service struct {
	Rows []*Row
	mux  sync.Mutex
}

func NewService(rows []*Row) *Service {
	return &Service{
		Rows: rows,
		mux:  sync.Mutex{},
	}
}

func (s *Service) getRow(uuid string) *Row {
	for _, v := range s.Rows {
		if v.UUID == uuid {
			return v
		}
	}
	return nil
}

// Delete 删除下载
func (s *Service) Delete(uuid *string, reply **Row) error {
	val := s.getRow(*uuid)
	if val == nil {
		return errors.New("uuid not find")
	}
	err := val.Delete()
	if err != nil {
		return err
	}
	*reply = val
	return nil
}

// Pause 暂停下载
func (s *Service) Pause(uuid *string, reply **Row) error {
	val := s.getRow(*uuid)
	if val == nil {
		return errors.New("uuid not find")
	}
	err := val.Pause()
	if err != nil {
		return err
	}
	*reply = val
	return nil
}

// Start 开始下载
func (s *Service) Start(uuid *string, reply **Row) error {
	val := s.getRow(*uuid)
	if val == nil {
		return errors.New("uuid not find")
	}
	err := val.Start()
	if err != nil {
		return err
	}
	*reply = val
	return nil
}

// GetRows 获取下载列表
func (s *Service) GetRows(args *struct{ Status string }, reply *[]*Row) error {
	statusVal := Status(args.Status)
	for _, v := range s.Rows {
		if statusVal != "" && !v.Status.Is(statusVal) {
			continue
		}

		if statusVal == "" && v.Status.Is(STATUS_COVERY) {
			continue
		}

		*reply = append(*reply, v)
	}
	return nil
}

// AddUri 新增资源下载
func (s *Service) AddUri(args *struct {
	URI     string
	Header  map[string]string
	Outdir  string
	Outname string
}, reply **Row) error {
	s.mux.Lock()

	uuid := NewUUID()
	row := &Row{
		UUID:   uuid,
		Status: STATUS_PAUSE,

		URI:     args.URI,
		Outdir:  args.Outdir,
		Outname: args.Outname,
		Header:  args.Header,

		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	s.Rows = append(s.Rows, row)

	s.mux.Unlock()

	err := row.Start()
	if err != nil {
		return err
	}

	*reply = row
	return nil
}

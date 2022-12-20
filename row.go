package main

import (
	"net/http"

	"github.com/rock-rabbit/rain"
)

type Row struct {
	Status   string
	UUID     string
	URI      string
	Progress int
	Outdir   string
	Outname  string
	Error    string
	Header   map[string]string
}

type RowEvent struct {
	*Row
}

func (r *RowEvent) Change(stat *rain.Stat) {
	r.Progress = stat.Progress

	if stat.Status == rain.STATUS_FINISH {
		if stat.Error != nil {
			r.SetError(stat.Error)
		} else {
			r.Status = STATUS_FINISH
		}
	}
}

func (r *Row) SetError(err error) {
	r.Status = STATUS_ERROR
	r.Error = err.Error()
}

func (r *Row) Start() {
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

	ctl := rain.New(r.URI, opts...)
	go func() {
		err := ctl.Start()
		if err != nil {
			r.SetError(err)
		}
	}()
}

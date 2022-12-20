package main

import (
	"encoding/json"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteControl struct {
	DB *gorm.DB
}

type RowTable struct {
	UUID   string `gorm:"primaryKey"`
	Status string

	URI     string
	Outdir  string
	Outname string
	Header  string

	Progress int
	Error    string

	CreateTime time.Time
	UpdateTime time.Time
}

func (rt *RowTable) Row() *Row {
	var header map[string]string
	json.Unmarshal([]byte(rt.Header), &header)
	return &Row{
		UUID:     rt.UUID,
		Status:   Status(rt.Status),
		URI:      rt.URI,
		Outdir:   rt.Outdir,
		Outname:  rt.Outname,
		Header:   header,
		Error:    rt.Error,
		Progress: rt.Progress,

		CreateTime: rt.CreateTime,
		UpdateTime: rt.UpdateTime,

		SyncTime: time.Now(),
	}
}

func NewSqliteControl() (*SqliteControl, error) {
	db, err := gorm.Open(sqlite.Open(config.Database), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&RowTable{})
	return &SqliteControl{DB: db}, nil
}

// GetRows 获取所有数据
func (ctl *SqliteControl) GetRows() ([]*RowTable, error) {
	var rows []*RowTable
	tx := ctl.DB.Find(&rows)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return rows, nil
}

// Update 更新数据库
func (ctl *SqliteControl) Update(rows []*Row) {
	for _, v := range rows {
		// 新建
		if v.SyncTime.Unix() < 0 {
			ctl.DB.Create(v.Table())
			v.SyncTime = time.Now()
			continue
		}

		// 更新
		if v.UpdateTime.Unix() > v.SyncTime.Unix() {
			ctl.DB.Save(v.Table())
			v.SyncTime = time.Now()
		}
	}

}

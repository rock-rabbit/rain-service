package main

import (
	"os"
	"path/filepath"

	uuid "github.com/satori/go.uuid"
)

func NewUUID() string {
	return uuid.NewV4().String()
}

// GetExecutable 获取执行文件目录
func GetExecutable() string {
	p, err := os.Executable()
	if err != nil {
		p, _ = filepath.Abs(filepath.Dir(os.Args[0]))
		return p
	}
	return filepath.Dir(p)
}

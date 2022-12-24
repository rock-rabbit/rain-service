package main

import "path/filepath"

type Config struct {
	Database string
}

var config = &Config{
	Database: filepath.Join(GetExecutable(), "rain-service.db"),
}

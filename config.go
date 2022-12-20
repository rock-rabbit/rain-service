package main

type Config struct {
	Database string
}

var config = &Config{
	Database: "./rain-service.db",
}

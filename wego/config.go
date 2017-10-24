package wego

import (
	"time"
)

func NewConfig() *WegoConfig {
	w := NewWebConfig()
	return &WegoConfig{WebConfig: w}
}

func NewWebConfig() *WegoWebConfig {
	return &WegoWebConfig{
		Addr:         ":8080",
		ReadTimeout:  time.Duration(10) * time.Second,
		WriteTimeout: time.Duration(10) * time.Second,
	}
}

type WegoConfig struct {
	WebConfig *WegoWebConfig
}

type WegoWebConfig struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

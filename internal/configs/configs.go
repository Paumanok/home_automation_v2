package configs

import (
	"time"
	"datapaddock.lan/go_server/internal/server/http"
)

type Configs struct {
}


func (cfg *Configs) HTTP() (*http.Config, error) {
	return &http.Config{
		Host: "localhost",
		Port: "8080",
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}, nil
}

// New returns an instance of Config with all the required dependencies initialized
func New() (*Configs, error) {
	return &Configs{}, nil
}
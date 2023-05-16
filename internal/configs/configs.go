package configs

import (
	"time"
	"datapaddock.lan/go_server/internal/server/http"
	"datapaddock.lan/go_server/internal/utils/database"
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

func (cfg  *Configs) Database() (*database.Config, error) {
	return &database.Config{
		Host: 	"postgres",
		Port:   "5424",
		Driver: "postgres",
		
		SSLMode: "",		

		StoreName: "HomeAuto",
		Username:  "user",
		Password:  "password",

		ConnPoolSize: 10,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		IdleTimeout:  time.Second * 60,
		DialTimeout:  time.Second * 10,
	}, nil
}

// New returns an instance of Config with all the required dependencies initialized
func New() (*Configs, error) {
	return &Configs{}, nil
}
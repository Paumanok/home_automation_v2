package configs

import (
	"time"
	"io/ioutil"
	"fmt"
	"datapaddock.lan/go_server/internal/server/http"
	"datapaddock.lan/go_server/internal/utils/database"
	"gopkg.in/yaml.v3"
)

type Configs struct {
	HttpCfg *http.Config `yaml:"http"`
	DatabaseCfg *database.Config `yaml:"database"`
}


func (cfg *Configs) HTTP() (*http.Config, error) {
	return &http.Config{
		Host: "0.0.0.0",
		Port: "8080",
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}, nil
}

func (cfg *Configs) ReadConfigFile() error {
	yamlFile, err := ioutil.ReadFile("cfg.yml")
	if err != nil {
		return err
	}

	var config = new(Configs)

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println(err)
		return err
	}
	
	cfg.HttpCfg = config.HttpCfg
	cfg.DatabaseCfg = config.DatabaseCfg
	return nil
}

func (cfg  *Configs) Database() (*database.Config, error) {
	return &database.Config{
		Host: 	"postgres",
		Port:   "5432",
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
	var cf = new(Configs)
	err := cf.ReadConfigFile()
	if err != nil {
		fmt.Println("read config failed")
	}
	return cf, nil	
}

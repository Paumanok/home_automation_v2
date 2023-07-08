package main

import (
	"fmt"
	"os"
	"datapaddock.lan/go_server/internal/server/http"
	"datapaddock.lan/go_server/internal/utils/database"
	"datapaddock.lan/go_server/internal/configs"
	"datapaddock.lan/go_server/internal/devices"
	"datapaddock.lan/go_server/internal/measurements"
)

func main() {
	fmt.Println("Hello World")
	var databaseCfg = new(database.Config)
	var httpCfg = new(http.Config)
	cfg, err := configs.New()
	if err != nil {
		//use hardcoded cfg

		databaseCfg, err = cfg.Database()
		if err != nil {
			fmt.Println("db config creation failed")
		}

		httpCfg, err = cfg.HTTP()
		if err != nil {
			fmt.Println("config creation failed")
			os.Exit(0)
		}
	} else {
		httpCfg = cfg.HttpCfg
		databaseCfg = cfg.DatabaseCfg
	}

	pool, err := database.NewService(databaseCfg)
	if err != nil {
		fmt.Println("pgxpool start failed")
	} 

	deviceStore, err := devices.NewStore(pool)
	if err != nil {
		fmt.Println("deviceStore failed")
	}
	
	deviceService, err := devices.NewService(deviceStore)
	if err != nil {
		fmt.Println("deviceService failed")
	}

	measurementStore, err := measurements.NewStore(pool)
	if err != nil {
		fmt.Println("measurementStore failed")
	}

	measurementService, err := measurements.NewService(measurementStore)
	if err != nil {
		fmt.Println("measurementService failed")
	}


	http_server, err := http.NewService(httpCfg, measurementService, deviceService)
	if err != nil {
		fmt.Println("http server start failed")
		os.Exit(0)
	}
	
	http_server.Start()
}

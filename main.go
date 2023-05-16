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
	cfg, err := configs.New()

	databaseCfg, err := cfg.Database()
	if err != nil {
		fmt.Println("db config creation failed")
	}

	pool, err := database.NewService(databaseCfg)
	if err != nil {
		fmt.Println("pgxpool start failed")
	} 

	deviceStore, err := devices.NewStore(pool)
	if err != nil {
		fmt.Println("deviceStore failed")
	}
	
	measurementStore, err := measurements.NewStore(pool)
	if err != nil {
		fmt.Println("measurementStore failed")
	}

	httpCfg, err := cfg.HTTP()
	if err != nil {
		fmt.Println("config creation failed")
		os.Exit(0)
	}

	http_server, err := http.NewService(httpCfg)
	if err != nil {
		fmt.Println("http server start failed")
		os.Exit(0)
	}
	
	http_server.Start()
}

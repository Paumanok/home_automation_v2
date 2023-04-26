package main

import (
	"fmt"
	"os"
	"datapaddock.lan/go_server/internal/server/http"
	"datapaddock.lan/go_server/internal/configs"
)

func main() {
	fmt.Println("Hello World")
	cfg, err := configs.New()

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

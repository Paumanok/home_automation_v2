package http

import (
	"fmt"
	"net/http"

	"time"

)

type HTTP struct {
	server *http.Server
	cfg    *Config
}

type Config struct {
	Host string
	Port string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

func (h *HTTP) Start() error {
	return h.server.ListenAndServe()
}



func NewService(cfg *Config) (*HTTP, error) {
	fmt.Println("Hello World from http service")
	//t := template.New("web/templates/index.html")

	//home, err := t.ParseFiles("web/templates/index.html",)
	//if err != nil {
	//	return nil, err
	//}
	
	baseHandler := &BaseHandler{
		MeasurementHandler: new(MeasurementHandler),
	}

	httpServer := &http.Server{
		Addr:		fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Handler:	baseHandler, 
		ReadTimeout: cfg.ReadTimeout, 
		WriteTimeout: cfg.WriteTimeout,
	}

	
	return &HTTP{ 
		server: httpServer,
		cfg: cfg,
	}, nil
}
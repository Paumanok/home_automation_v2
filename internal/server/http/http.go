package http

import (
	"fmt"
	"net/http"

	"time"
	"datapaddock.lan/go_server/internal/measurements"
	//"datapaddock.lan/go_server/internal/devices"
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



func NewService(cfg *Config, m *measurements.Measurements) (*HTTP, error) {
	fmt.Println("Hello World from http service")
	//t := template.New("web/templates/index.html")

	//home, err := t.ParseFiles("web/templates/index.html",)
	//if err != nil {
	//	return nil, err
	//}
	
	baseHandler := &BaseHandler{
		MeasurementHandler: &MeasurementHandler{
			m : m,
		},
		IndexHandler: new(IndexHandler),
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
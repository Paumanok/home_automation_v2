package http

import (
	"fmt"
	"net/http"

	"time"
	"datapaddock.lan/go_server/internal/measurements"
	"datapaddock.lan/go_server/internal/devices"
	"datapaddock.lan/go_server/internal/utils/helpers"
)

type HTTP struct {
	server *http.Server
	cfg    *Config
}

type Config struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
	ReadTimeout time.Duration `yaml:"ReadTimeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
}

func (h *HTTP) Start() error {
	return h.server.ListenAndServe()
}

// CORS middleware
func withCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://kasa.datapaddock.lan") // Allow all origins, adjust for security as needed
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	fmt.Println("with cors?")
        // Handle preflight requests
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
	    fmt.Println("method okay")
            return
        }

        // Call the next handler
        next.ServeHTTP(w, r)
    })
}


func NewService(cfg *Config, m *measurements.Measurements, d *devices.Devices) (*HTTP, error) {
	fmt.Println("Hello World from http service")
	//t := template.New("web/templates/index.html")

	//home, err := t.ParseFiles("web/templates/index.html",)
	//if err != nil {:
	//	return nil, err
	//}
	
	measurement_handler := &MeasurementHandler{
		service: m,
		base: nil,
	}

	device_handler := &DeviceHandler{
		service: d,
		base: nil,
	}

	sync_timer := &helpers.SyncTimer{
		TimerInterval: 60,
		TimerVal: 60,
	}

	baseHandler := &BaseHandler{
		MeasurementHandler: measurement_handler,
		DeviceHandler: device_handler,
		IndexHandler: new(IndexHandler),
		SyncTimer: sync_timer,
	}

	measurement_handler.base = baseHandler
	device_handler.base = baseHandler	

	handler := withCORS(baseHandler)

	go sync_timer.Timer()

	httpServer := &http.Server{
		Addr:		fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Handler:	handler, 
		ReadTimeout: cfg.ReadTimeout, 
		WriteTimeout: cfg.WriteTimeout,
	}

	
	return &HTTP{ 
		server: httpServer,
		cfg: cfg,
	}, nil
}

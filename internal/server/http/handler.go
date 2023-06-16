package http

import (
	"fmt"
	"net/http"
	//"strings"
	"encoding/json"
	"html/template"
	"datapaddock.lan/go_server/internal/utils/helpers"
	"path/filepath"
	"datapaddock.lan/go_server/internal/measurements"
	"datapaddock.lan/go_server/internal/devices"

)
//empty structs take up no space but enable it
//to be used as a "method receiver"
// I think the empty struct inside of it might also make this zero bytes but idk
type BaseHandler struct {
	MeasurementHandler *MeasurementHandler
	DeviceHandler *DeviceHandler
	IndexHandler *IndexHandler
	SyncTimer *helpers.SyncTimer
}	
						


func (h *BaseHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = helpers.ShiftPath(req.URL.Path)
	//here we're going to define our basic api endpoints
	// i'm going to supply endpoints for the old server for now until I refactor
	// the esp code

	switch head {
	case "":
		fmt.Println("hit index case")
		h.IndexHandler.ServeHTTP(res, req)
		//index/homepage
	case "images":
		//return jpeg of plot
	case "next":
		h.ServeNext(res, req)	
		//get next measurement delay
	case "config":
		//configuration page
	case "measurements":
		//root of measurements queries
		//subsequent ones will check for /measurements/last or /measurements/start:end
		//new measurements from devices should also eventually go through here
		h.MeasurementHandler.ServeHTTP(res, req)
	case "data":
		//device data POST endpoint
		h.ServeData(res, req)
	default:
		fmt.Println("hit default")
	}
}

func (h *BaseHandler) ServeNext(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		st_resp := fmt.Sprint("sync_time ", h.SyncTimer.GetNextDelay())
		res.Header().Set("Content-Type", "text/plain")
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(st_resp))
	}
}

func (h *BaseHandler) ServeData(res http.ResponseWriter, req *http.Request){
	
	switch req.Method {
	case "POST":
		measurement := new(measurements.Measurement)
		err := json.NewDecoder(req.Body).Decode(measurement)
		if err != nil {
			fmt.Println("json decode failed")
			return
		}
		exists, err := h.DeviceHandler.service.DeviceExists(req.Context(), measurement.MAC)
		if err != nil {
			fmt.Println("error checking for device existance")
		}
		
		if !exists {
			fmt.Println("device does not exist, adding")
			var d = new(devices.Device)
			d.MAC = measurement.MAC
			d, err :=h.DeviceHandler.service.CreateDevice(req.Context(), d)
			if err != nil {
				fmt.Println(err)
				return
			}

		}

		exists, err = h.DeviceHandler.service.DeviceExists(req.Context(), measurement.MAC)
		if err != nil {
			fmt.Println("error checking for device existance")
		}
		
		if exists {
			h.MeasurementHandler.service.CreateMeasurement(req.Context(), measurement)
		}else{
			fmt.Println("why doesn't it exist")
		}

		return

	case "GET":
		//not sure if we need something here
	}
}

type MeasurementHandler struct {
	service *measurements.Measurements
	base *BaseHandler
}

func (h *MeasurementHandler) ServeHTTP(res http.ResponseWriter, req *http.Request){

	switch req.Method {
	case "POST":
		//getting data from esp
		measurement := new(measurements.Measurement)
		err := json.NewDecoder(req.Body).Decode(measurement)
		if err != nil {
			fmt.Println("json decode failed")
			return
		}
		fmt.Println("handler creating measurement")
		h.service.CreateMeasurement(req.Context(), measurement)
		return

	case "GET":
		var meas  []measurements.Measurement
		var err error
		var head string
		head, req.URL.Path = helpers.ShiftPath(req.URL.Path)
		fmt.Println(head)
		//getting data from db
		fmt.Println("in get")
		switch head {
		case "":
			meas, err = h.service.GetAllMeasurements(req.Context())
		
		case "last":
			meas, err = h.ServeLast(res, req)
		
		case "range":
			meas, err = h.ServePeriod(res, req)
		}

		if err != nil {
			fmt.Println(err)
			return
		}
		
		if meas == nil {
			fmt.Println("No Measurements found")
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(meas)

	}
}

type DeviceHandler struct {
	service *devices.Devices
	base *BaseHandler
}

type IndexHandler struct {}

func (h *IndexHandler) ServeHTTP(res http.ResponseWriter, req *http.Request){
	fmt.Println("index called")

	template_path, err := filepath.Abs("internal/server/http/web/templates/")
	if err != nil {
		fmt.Println(err)
		return
	}
	index_path := template_path + "/" + "index.html"
	
	t := template.New(index_path)

	t, err = t.ParseFiles(index_path)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := struct {
		Name string
	}{"myname"}

	t.ExecuteTemplate(res,"index.html", data)
	return 
}

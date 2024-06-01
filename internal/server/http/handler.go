package http

import (
	"fmt"
	"net/http"
	"strings"
	"encoding/json"
	//"html/template"
	"datapaddock.lan/go_server/internal/utils/helpers"
	//"path/filepath"
	"datapaddock.lan/go_server/internal/measurements"
	"datapaddock.lan/go_server/internal/devices"
	"datapaddock.lan/go_server/internal/server/frontend"

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
	
	case "api":
		h.ServeApi(res, req)
	default:
		fmt.Println(req.URL.Path)
		fmt.Println("hit default")
		req.URL.Path = head + req.URL.Path
		h.IndexHandler.ServeHTTP(res, req)
	}
}

func (h *BaseHandler) ServeApi(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = helpers.ShiftPath(req.URL.Path)
	switch head {

	case "images":
		//return jpeg of plot
	case "next":
		h.ServeNext(res, req)	
		//get next measurement delay
	case "config":
		//configuration page
	case "devices":
		h.DeviceHandler.ServeHTTP(res, req)
	case "measurements":
		//root of measurements queries
		//subsequent ones will check for /measurements/last or /measurements/start:end
		//new measurements from devices should also eventually go through here
		h.MeasurementHandler.ServeHTTP(res, req)
	case "data":
		//device data POST endpoint
		h.ServeData(res, req)
	default:
		fmt.Println("api default")
	}
}

func (h *BaseHandler) ServeNext(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		//fmt.Println(req.Header)
		//fmt.Println("NEXT HIT")
		if req.Header.Get("Accept") == "text/plain" {
			st_resp := fmt.Sprint("sync_time ", h.SyncTimer.GetNextDelay())
			res.Header().Set("Content-Type", "text/plain")
			res.WriteHeader(http.StatusOK)
			res.Write([]byte(st_resp))
		} else if req.Header.Get("Accept") == "application/json" {
			//doing this naevely because it's only a single key/value
			st_resp := fmt.Sprint("{\"sync_time\":", h.SyncTimer.GetNextDelay(), "}")
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusOK)
			res.Write([]byte(st_resp))
		}
	}
}

//backwards compatible device data handler, original design had devices hitting a /data endpoint to 
//shove json into, but we've evolved.
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
		fmt.Println("handler creating measurement")
		h.CollectDataFromDevice(res, req)
		return

	case "GET":
		//var meas  []measurements.Measurement
		var meas any
		var err error
		var head string
		head, req.URL.Path = helpers.ShiftPath(req.URL.Path)
		//getting data from db
		fmt.Printf("MeasurementHandler:ServeHTTP---Head: %s\n", head )

		switch head {
		case "":
			meas, err = h.service.GetAllMeasurements(req.Context())
		
		case "last":
			fmt.Println("do we see this?")
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

func (h *DeviceHandler) ServeHTTP(res http.ResponseWriter, req *http.Request){
	switch req.Method {
	case "GET":
		var devs []devices.Device
		var err error
		var head string
		head, req.URL.Path = helpers.ShiftPath(req.URL.Path)
		
		switch head {
		case "":
			devs, err = h.service.GetAllDevices(req.Context())
		case "update":
			dev, err := h.updateDevice(res, req)
			if err != nil {
				fmt.Println(err)
			} else {
				devs = append(devs, *dev)
			}
		default:
			//nothing here 	
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(devs)
	}	
}

type IndexHandler struct {}

func (h *IndexHandler) ServeHTTP(res http.ResponseWriter, req *http.Request){
	fmt.Println("index called")
	fa := frontend.GetFrontendAssets()
	//naive mimetype stuff
	if strings.HasSuffix(req.URL.Path, ".css"){
		res.Header().Set("Content-Type","text/css")
	} else if strings.HasSuffix(req.URL.Path, ".html") {
		res.Header().Set("Content-Type", "text/html")
	}else if strings.HasSuffix(req.URL.Path, ".js") {
		res.Header().Set("Content-Type", "text/javascript")
	}

	http.FileServer(http.FS(fa)).ServeHTTP(res, req)
	//template_path, err := filepath.Abs("internal/server/http/web/templates/")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//index_path := template_path + "/" + "index.html"
	//
	//t := template.New(index_path)

	//t, err = t.ParseFiles(index_path)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//data := struct {
	//	Name string
	//}{"myname"}

	//t.ExecuteTemplate(res,"index.html", data)
	return 
}

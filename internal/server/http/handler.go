package http

import (
	"fmt"
	"net/http"
	"encoding/json"
	"html/template"
	"datapaddock.lan/go_server/internal/utils/helpers"
	"path/filepath"
	"datapaddock.lan/go_server/internal/measurements"
	//"datapaddock.lan/go_server/internal/devices"

)
//empty structs take up no space but enable it
//to be used as a "method receiver"
// I think the empty struct inside of it might also make this zero bytes but idk
type BaseHandler struct {
	MeasurementHandler *MeasurementHandler
	IndexHandler *IndexHandler
}	
						


func (h *BaseHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = helpers.ShiftPath(req.URL.Path)
	fmt.Println("hit base handler")
	//here we're going to define our basic api endpoints
	// i'm going to supply endpoints for the old server for now until I refactor
	// the esp code
	fmt.Println(head)
	switch head {
	case "":
		fmt.Println("hit index case")
		h.IndexHandler.ServeHTTP(res, req)
		//index/homepage
	case "images":
		//return jpeg of plot
	case "next":
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
	default:
		fmt.Println("hit default")
	}
}

type MeasurementHandler struct {
	service *measurements.Measurements
}

func (h *MeasurementHandler) ServeHTTP(res http.ResponseWriter, req *http.Request){
	var head string
	head, req.URL.Path = helpers.ShiftPath(req.URL.Path)
	fmt.Println(head)
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
		//getting data from db
	}
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

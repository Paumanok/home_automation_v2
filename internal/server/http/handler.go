package http

import (
	"fmt"
	"net/http"
	"html/template"
	"datapaddock.lan/go_server/internal/utils/helpers"
)
//empty structs take up no space but enable it
//to be used as a "method receiver"
// I think the empty struct inside of it might also make this zero bytes but idk
type BaseHandler struct {
	MeasurementHandler *MeasurementHandler
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
		index(res, req)
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

type MeasurementHandler struct {}

func (h *MeasurementHandler) ServeHTTP(res http.ResponseWriter, req *http.Request){
	var head string
	head, req.URL.Path = helpers.ShiftPath(req.URL.Path)
	fmt.Println(head)
	switch req.Method {
	case "POST":
		//getting data from esp
	case "GET":
		//getting data from db
	}
}

func index(res http.ResponseWriter, req *http.Request){
	fmt.Println("index called")
	t := template.New("internal/server/web/templates/index.html")

	t, err := t.ParseFiles("internal/server/web/templates/index.html",)
	if err != nil {
		fmt.Println("error?")
		fmt.Println(err)
		return
	}

	t.Execute(res, nil)
}

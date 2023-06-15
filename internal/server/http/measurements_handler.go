package http

import (
	"fmt"
	"net/http"
	"time"
	"datapaddock.lan/go_server/internal/utils/helpers"
	"datapaddock.lan/go_server/internal/measurements"
	//"datapaddock.lan/go_server/internal/devices"
)


func (h *MeasurementHandler) ServeLast(res http.ResponseWriter, req *http.Request) ([]measurements.Measurement, error){
	//dh := h.base.DeviceHandler
	var meas []measurements.Measurement
	var err error
	var head string
	
	if req.Method != "GET" {
		return nil, nil
	}

	head, req.URL.Path = helpers.ShiftPath(req.URL.Path)
	
	switch head {

	case "":
		macs := h.GetMacs(req)
		interval := h.base.SyncTimer.TimerInterval
		meas, err = h.service.GetLastMeasurements(req.Context(), macs, interval )
	
	case "hour":
		interval := 60*60
		now := time.Now()
		cutoff := now.Add(time.Duration(-1*interval)*time.Second)
		meas, err = h.service.GetMeasurementsSince(req.Context(), cutoff)

	case "day":
		meas, err = h.service.GetLastNumDays(req.Context(), 1)

	case "week":
		meas, err = h.service.GetLastNumDays(req.Context(), 7)
	
	case "month":
		meas, err = h.service.GetLastNumDays(req.Context(), 30)
	}

	
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return meas, nil

}



func (h *MeasurementHandler) GetMacs(req *http.Request) []string {
	dh := h.base.DeviceHandler
	var macs []string
	devs,_ := dh.service.GetAllDevices(req.Context())
	for _, dev := range devs {
		macs = append(macs, dev.MAC)
	}

	return macs
}

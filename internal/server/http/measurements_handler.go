package http

import (
	"fmt"
	"net/http"
	"net/url"
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

//this handles something looking like this http://192.168.0.151:8080/measurements/range?end=2023-06-16T01%3A04%3A46.814883%2B00%3A00&start=2023-06-15T01%3A04%3A46.814883%2B00%3A00
//datetime must be url-encoded, requires timezone
//python to generate: urllib.parse.quote(datetime.datetime.utcnow().replace(tzinfo=datetime.timezone.utc).isoformat())
//this comment should go away eventually
func (h *MeasurementHandler) ServePeriod(res http.ResponseWriter, req *http.Request) ([]measurements.Measurement, error) {
	fmt.Println(req.URL.RawQuery)
	q , err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(q["start"])
	fmt.Println(q["end"])
	if !q.Has("start") || !q.Has("end") {
		fmt.Println("not enough params")
		return nil, nil
	}

	start, err := time.Parse(time.RFC3339, q["start"][0])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	end, err := time.Parse(time.RFC3339, q["end"][0])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(start)
	fmt.Println(end)
	meas, err := h.service.GetMeasurementsFromPeriod(req.Context(), start, end)
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

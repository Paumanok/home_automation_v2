package http

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
	"encoding/json"
	//"datapaddock.lan/go_server/internal/utils/helpers"
	"datapaddock.lan/go_server/internal/measurements"
	"datapaddock.lan/go_server/internal/devices"
)


// Create new measurement, ensure device exists in db for said measurement, 
// if it doesn't exist, add it
func (h *MeasurementHandler) CollectDataFromDevice( res http.ResponseWriter, req *http.Request) {
	measurement := new(measurements.Measurement)
	err := json.NewDecoder(req.Body).Decode(measurement)
	if err != nil  {
		fmt.Println("json decode failed")
		return
	}
	exists, err := h.base.DeviceHandler.service.DeviceExists(req.Context(), measurement.MAC)
	if err != nil {
		fmt.Println("error checking for device existance")
	}
	if !exists {
		fmt.Println("device does not exist, adding")
		var d = new(devices.Device)
		d.MAC = measurement.MAC
		d, err :=h.base.DeviceHandler.service.CreateDevice(req.Context(), d)
		if err != nil {
			fmt.Println(err)
			return
		}

	}

	exists, err = h.base.DeviceHandler.service.DeviceExists(req.Context(), measurement.MAC)
	if err != nil {
		fmt.Println("error checking for device existance")
	}
	
	if exists {
		h.service.CreateMeasurement(req.Context(), measurement)
	}else{
		fmt.Println("why doesn't it exist")
	}
	return
}

// So I don't forget a second time,
// I had converted ServeLast to use query paramters
// ex: /api/measurements/last?period=day&byDevice=true gets the last day of measurements sorted by device
func (h *MeasurementHandler) ServeLast(res http.ResponseWriter, req *http.Request) (any, error){
	//dh := h.base.DeviceHandler
	var meas []measurements.Measurement
	var err error
	var head string

	//fmt.Println(req.URL.RawQuery)
	q , err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	
	if req.Method != "GET" {
		return nil, nil
	}
	if !q.Has("period") {
		return nil, nil
	}
	//head, req.URL.Path = helpers.ShiftPath(req.URL.Path)
	head = q["period"][0]	
	switch head {

	case "last":
		macs := h.GetMacs(req)
		interval := h.base.SyncTimer.TimerInterval
		fmt.Println(interval)
		meas, err = h.service.GetLastMeasurements(req.Context(), macs, interval )
		if err != nil {
			fmt.Println(err)
		}
	
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
	if q.Has("fahrenheit") {
		if q["fahrenheit"][0] == "true" {
			fmt.Println("c2f")
			meas = h.c2f(meas)
		}

	}
	if q.Has("comp") {
		if q["comp"][0] == "true" {
			meas = h.compensateMeasurements(req, meas)
		}
	}
	if q.Has("byDevice") {
		if q["byDevice"][0] == "true" {
			fmt.Println("hit")
			sortmeas, err := h.sortMeasurements(req, meas)
			if err == nil {
				return sortmeas, nil
			}
		}
	}

	return meas, nil

}

func (h *MeasurementHandler) c2f(meas []measurements.Measurement) ([]measurements.Measurement) {
	for idx, measurement := range meas {
		measurement.Temp = (measurement.Temp * 9/5) + 32
		meas[idx] = measurement
	}
	return meas
}

func (h *MeasurementHandler) compensateMeasurements(req *http.Request, meas []measurements.Measurement) ([]measurements.Measurement) {
	var devs = make(map[string]devices.Device)
	macs := h.GetMacs(req)

	for _, mac := range macs {
		dev, _ := h.base.DeviceHandler.service.GetDeviceByMac(req.Context(), mac)
		devs[mac] = *dev

	}

	for idx, measurement := range meas {
		measurement.Temp += float32(devs[measurement.MAC].TemperatureComp)
		measurement.Humidity += float32(devs[measurement.MAC].HumidityComp)
		meas[idx] = measurement
	}

	return meas
}

type SortedMeasurement struct {
	DeviceInfo devices.Device `json:"deviceInfo"`
	Measurements []measurements.Measurement `json:"measurements"`
}

func (h *MeasurementHandler) sortMeasurements(req *http.Request, meas []measurements.Measurement) (map[string]SortedMeasurement, error) {
	var sorted = make(map[string]SortedMeasurement)
	
	for _, measurement := range meas {
		mac := measurement.MAC
		bucket, ok := sorted[mac]
		if ok {
			bucket.Measurements = append(bucket.Measurements, measurement)
			sorted[mac] = bucket
		} else {
			var m []measurements.Measurement
			m = append(m, measurement)
			device_info, err := h.base.DeviceHandler.service.GetDeviceByMac(req.Context(), mac)
			if err != nil {
				return nil, err
			}
			sorted[mac] = SortedMeasurement {
				DeviceInfo: *device_info, 
				Measurements: m,
			}
		}
	}

	return sorted, nil
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

package http

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"datapaddock.lan/go_server/internal/devices"
)


func (h *DeviceHandler) updateDevice(res http.ResponseWriter, req *http.Request) (*devices.Device, error){
	var err error
	var mac string

	q, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if req.Method != "GET" {
		return nil, nil
	}

	if !q.Has("mac") {
		return nil, nil
	}

	mac = q["mac"][0]

	// since we have the mac, lets ensure it exists
	exists, err := h.service.DeviceExists(req.Context(), mac)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if !exists {
		fmt.Println("Cannot update device that does not exist")
		return nil, nil
	}
	// lets get the existing device from the db since we know it exists
	device, err := h.service.GetDeviceByMac(req.Context(), mac)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	//lets modify according to props
	if q.Has("nickname") {
		var nickname = q["nickname"][0]
		//TODO: add some limitations to this, for now, just function
		device.Nickname = nickname
	}
	if q.Has("humidity_comp") {
		h_comp, err := strconv.Atoi(q["humidity_comp"][0])
		if err != nil {
			fmt.Println(err)
			fmt.Println("invalid value for humidity_comp")
		} else {
			device.HumidityComp = h_comp
		}
	}
	if q.Has("temp_comp") {
		temp_comp, err := strconv.Atoi(q["temp_comp"][0])
		if err != nil {
			fmt.Println(err)
			fmt.Println("invalid value for temp_comp")
		} else {
			device.TemperatureComp = temp_comp
		}
	}

	device, err = h.service.UpdateDevice(req.Context(), device)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to update device")
	}

	return device, nil
}

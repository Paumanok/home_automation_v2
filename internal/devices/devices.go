package devices

import (
	"fmt"
	"context"
	"strings"
	"strconv"
)


type Device struct {
	Nickname string `json:"nickname"`
	MAC		 string `json:"mac"`
	HumidityComp int `json:"humidityComp"`
	TemperatureComp int `json:"temperatureComp"`
}

type store interface {
	Create(ctx context.Context, d *Device) error
	Update(ctx context.Context, d *Device, mac string) error
	GetDeviceByMac(ctx context.Context, mac string) (*Device, error)
	GetDevices(ctx context.Context) ([]Device, error)
}

type Devices struct {
	store store
}

func (ds *Devices) CreateDevice(ctx context.Context, d *Device) (*Device, error) {
	fmt.Println("Creating device from devices.go")
	if !ds.validateMac(d.MAC) {
		fmt.Println("invalid MAC,")
		return nil, nil
	}
	
	if len(d.Nickname) == 0 {
		d.Nickname = d.MAC
	}
	err := ds.store.Create(ctx, d)
	if err != nil {
		fmt.Println("db create device failed")
		return nil, err
	}
	return d, nil
}

func (ds *Devices) UpdateDevice(ctx context.Context, d *Device) (*Device, error) {
	fmt.Println("Updating device")

	exists, err := ds.DeviceExists(ctx, d.MAC)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if exists {
		err = ds.store.Update(ctx, d, d.MAC)
		if err != nil {
			fmt.Println("db update device failed")
			return nil, err
		}
	} else {
		fmt.Println("cannot update device, does not exist")
	}
	
	return d, nil
}

func (ds *Devices) GetAllDevices(ctx context.Context) ([]Device, error){
	var devs []Device
	devs, err := ds.store.GetDevices(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return devs, nil
}

func (ds *Devices) validateMac(mac string) bool {
	if len(mac) == 0 {
		return false
	}
	if strings.ContainsAny(mac, ":") {
		
		pairs := strings.Split(mac, ":")
		if len(pairs) != 6 {
			fmt.Println("length bad %d\n", len(pairs))
			return false
		}
		
		for _, val := range pairs {
			val, err := strconv.ParseInt(strings.ToLower(val), 16, 16)
			if err != nil {
				fmt.Println("parse err")
				fmt.Println(err)
			}
			if err != nil || val > 0xff {
				return false
			}
		}
	}
	return true
}

func (ds *Devices) GetDeviceByMac(ctx context.Context, mac string) (*Device, error) {
	if !ds.validateMac(mac) {
		return nil, nil
	}

	dev, err := ds.store.GetDeviceByMac(ctx, mac)

	return dev, err
}

func (ds *Devices) DeviceExists(ctx context.Context, mac string ) (bool, error) {
	if !ds.validateMac(mac) {
		fmt.Println("Mac not valid")
		return false, nil
	}
	dev, err := ds.store.GetDeviceByMac(ctx, mac)
	if dev == nil || err != nil{
		fmt.Println("in deviceExists, device does not, in fact, exist")
		return false, nil
	}
	return true, nil
}

func NewService(s store) (*Devices, error) {
	return &Devices {
		store: s,
	},nil
}

package devices

import (
	"fmt"
	"context"
)


type Device struct {
	Nickname string `json:"nickname"`
	MAC		 string `json:"mac"`
	HumidityComp int `json:"humidityComp"`
	TemperatureComp int `json:"temperatureComp"`
}

type store interface {
	Create(ctx context.Context, d *Device) error
	GetdeviceByMac(ctx context.Context, mac string) (*Device, error)
	GetDevices(ctx context.Context) ([]Device, error)
}

type Devices struct {
	store store
}

func (ds *Devices) CreateDevice(ctx context.Context, d *Device) (*Device, error) {
	fmt.Println("Creating device from devices.go")
	if len(d.MAC) == 0 {
		fmt.Println("invalid device, no MAC found")
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

func (ds *Devices) GetAllDevices(ctx context.Context) ([]Device, error){
	return nil, nil
}

func (ds *Devices) GetDeviceByMac(ctx context.Context, mac string) (*Device, error) {
	return nil, nil
}

func (ds *Devices) GetDeviceExists(ctx context.Context, mac string ) (bool, error) {
	return false, nil
}

func NewService(s store) (*Devices, error) {
	return &Devices {
		store: s,
	},nil
}
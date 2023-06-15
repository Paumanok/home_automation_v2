package measurements

import (
	"time"
	"fmt"
	"context"
)


type Measurement struct {
	MAC string `json:"MAC"`
	Temp float32 `json:"temp"`
	Humidity float32 `json:"humidity"`
	Pressure float32 `json:"pressure,omitempty"`
	PM25	float32  `json:"pm25,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}



type store interface {
	Create(ctx context.Context, m *Measurement) error
	GetByMAC(ctx context.Context, mac string) ([]Measurement, error)
	GetAllMeasurements(ctx context.Context) ([]Measurement, error)
	GetLastByMac(ctx context.Context, mac string) (*Measurement, error)
	GetSince(ctx context.Context, cutoff time.Time) ([]Measurement, error)
}

type Measurements struct {
	store store
}

func (ms *Measurements) CreateMeasurement( ctx context.Context, m *Measurement) (*Measurement, error) {
	
	//dear future me, write some functions to make sure the data is good. 
	if len(m.MAC) == 0 {
		fmt.Println("invalid measurement, no MAC found")
		//actually set up errors
		return nil, nil
	}

	if m.Temp == 0 || m.Humidity == 0 {
		fmt.Println("invalid measurement, no t or h")
		return nil,nil
	}
	
	if m.CreatedAt == nil {
		now := time.Now()
		m.CreatedAt = &now
	}
	fmt.Printf("Creating Measurement: %v+\n", m)
	err := ms.store.Create(ctx, m)
	if err != nil {
		fmt.Println("db create failed")
		return nil, err
	}
	return m, nil
}

func (ms *Measurements) GetMeasurementByMAC(ctx context.Context, MAC string) ([]Measurement, error){
	fmt.Printf("Getting measurements for MAC: %s", MAC)
	measurements, err := ms.store.GetByMAC(ctx, MAC)
	if err != nil {
		fmt.Println("failed to get measurements from db")
	}

	return measurements ,nil
}

func (ms *Measurements) GetAllMeasurements(ctx context.Context) ([]Measurement, error){
	fmt.Println("Getting all measurements")
	measurements, err := ms.store.GetAllMeasurements(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return measurements, nil
}
//This will only get the last measurement for a given device
//interval will not get more measurements if it's set longer, its merely a filter for
//how old a "last" measurement can be
func (ms *Measurements) GetLastMeasurements(ctx context.Context, macs []string, interval int) ([]Measurement, error) {
	var last_measurements []Measurement

	for _, mac := range macs {
		meas, err := ms.store.GetLastByMac(ctx, mac)
		if err != nil {
			fmt.Println(err)
			return nil, err 
		}
		now := time.Now()
		cutoff := now.Add(time.Duration(-1*interval)*time.Second)
		fmt.Println(cutoff)
		if meas.CreatedAt.After(cutoff) {
			last_measurements = append(last_measurements, *meas)
		}
	}
	return last_measurements, nil
}

func (ms *Measurements) GetLastNumDays(ctx context.Context, days int) ([]Measurement, error){
	meas, err := ms.GetMeasurementsSince(ctx, time.Now().Add(time.Duration(-days)*time.Hour))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return meas, nil
}

func (ms *Measurements) GetMeasurementsSince(ctx context.Context, cutoff time.Time) ([]Measurement, error) {
	meas, err := ms.store.GetSince(ctx, cutoff)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return meas,nil	
}

func NewService(s store) (*Measurements, error){
	return &Measurements{
		store: s,
	},nil
}

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
}

type Measurements struct {
	store store
}

func (ms *Measurements) CreateMeasurement( ctx context.Context, m *Measurement) (*Measurement, error) {
	
	//dear future me, write some functions to make sure the data is good. 
	fmt.Println("creating measurement from measurements.go")
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

func NewService(s store) (*Measurements, error){
	return &Measurements{
		store: s,
	},nil
}
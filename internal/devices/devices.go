package devices

import (
)


type Device struct {
	Nickname string `json:"nickname"`
	MAC		 string `json:"mac"`
	HumidityComp int `json:"humidityComp"`
	TemperatureComp int `json:"temperatureComp"`
}
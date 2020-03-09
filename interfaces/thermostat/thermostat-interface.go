package thermostat

type Thermostat interface {
	GetStatus()
	GetArrivalTemp()
	GetLeavingTemp()
	SetTemperature(temp int)
}

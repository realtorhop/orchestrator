package security_alarm

type SecurityAlarm interface {
	Arm()
	Disarm()
	GetStatus()
}

package lighting

type Lighting interface {
	On(room string)
	Off(room string)
	ListRooms()
	GetStatus(room string)
	ListLightsOn()
	ListLightsOff()
}

package door_lock

type DoorLock interface {
	Lock()
	Unlock()
	GetStatus()
}

package camera

type Camera interface {
	ListCameras()
	Record(camera string)
	Stop(camera string)
	View(camera string)
	Status()
}

package domain

type Box struct {
	ID     int
	Volume float64
	Weight float64
}

type Container struct {
	ID         int
	Boxes      []Box
	MaxWeight  float64
	MaxVolume  float64
	UsedWeight float64
	UsedVolume float64
}

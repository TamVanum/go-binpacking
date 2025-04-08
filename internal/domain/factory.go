package domain

func NewBigBox() *Box {
	return &Box{
		Volume: 30,
		Weight: 25,
	}
}

func NewMediumBox() *Box {
	return &Box{
		Volume: 20,
		Weight: 15,
	}
}

func NewSmallBox() *Box {
	return &Box{
		Volume: 10,
		Weight: 5,
	}
}

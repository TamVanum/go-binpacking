package domain

func NewBigBox(id int) *Box {
	return &Box{
		ID:     id,
		Volume: 30,
		Weight: 25,
	}
}

func NewMediumBox(id int) *Box {
	return &Box{
		ID:     id,
		Volume: 20,
		Weight: 15,
	}
}

func NewSmallBox(id int) *Box {
	return &Box{
		ID:     id,
		Volume: 10,
		Weight: 5,
	}
}

package domain

import "github.com/tamvanum/go-binpacking/pkg/utils"

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

func NewRandomDimensionBox() *Box {
	min, max := utils.RandomRangeFloat64()
	return &Box{
		Volume: min,
		Weight: max,
	}
}

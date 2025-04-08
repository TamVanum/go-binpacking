package main

import (
	"fmt"

	"github.com/tamvanum/go-binpacking/internal/domain"
)

func main() {
	boxes := []*domain.Box{
		domain.NewBigBox(1),
		domain.NewMediumBox(2),
		domain.NewSmallBox(3),
	}

	for _, b := range boxes {
		fmt.Printf("Box ID: %d | Volume: %.2f | Weight: %.2f\n", b.ID, b.Volume, b.Weight)
	}
}

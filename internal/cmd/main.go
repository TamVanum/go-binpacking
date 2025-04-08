package main

import (
	"fmt"

	"github.com/tamvanum/go-binpacking/internal/domain"
	"github.com/tamvanum/go-binpacking/internal/repository"
)

func main() {

	boxRepository := repository.NewInMemoryBoxRepository()

	boxRepository.Add(*domain.NewBigBox())
	boxRepository.Add(*domain.NewMediumBox())
	boxRepository.Add(*domain.NewMediumBox())
	boxRepository.Add(*domain.NewSmallBox())

	for _, b := range boxRepository.Get() {
		fmt.Printf("Box ID: %d | Volume: %.2f | Weight: %.2f\n", b.ID, b.Volume, b.Weight)
	}
}

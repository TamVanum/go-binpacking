package main

import (
	"fmt"

	"github.com/tamvanum/go-binpacking/internal/domain"
	"github.com/tamvanum/go-binpacking/internal/repository"
	"github.com/tamvanum/go-binpacking/internal/usecase"
)

func main() {

	boxRepository := repository.NewInMemoryBoxRepository()

	for range 5 {
		boxRepository.Add(*domain.NewBigBox())
		boxRepository.Add(*domain.NewMediumBox())
		boxRepository.Add(*domain.NewMediumBox())
		boxRepository.Add(*domain.NewSmallBox())
	}

	for _, b := range boxRepository.Get() {
		fmt.Printf("Box ID: %d | Volume: %.2f | Weight: %.2f\n", b.ID, b.Volume, b.Weight)
	}

	fmt.Println("-")
	fmt.Println("-")
	fmt.Println("-")

	firstContainer := domain.Container{ID: 1, MaxWeight: 100, MaxVolume: 100}

	containers := usecase.GreedyPack(boxRepository.Get(), firstContainer)

	for _, c := range containers {
		fmt.Printf("Container %d: -- WeightUesed %f | VolumeUsed %f | Boxes %d ", c.ID, c.UsedWeight, c.UsedVolume, len(c.Boxes))
		for _, b := range c.Boxes {
			fmt.Printf("Box %d", b.ID)
		}
		fmt.Println()
	}
}

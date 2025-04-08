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

	firstContainer := domain.Container{ID: 1, MaxWeight: 100, MaxVolume: 100}
	containerRepository := repository.NewContainerInMemoryRepository()

	for _, c := range usecase.GreedyPack(boxRepository.Get(), firstContainer) {
		containerRepository.Add(c)
	}

	for _, c := range containerRepository.Get() {
		fmt.Printf("ID %d | Leng %d | MaxVol %f | MaxWeight %f | UsedVol %f | UsedWeight %f \n", c.ID, len(c.Boxes), c.MaxVolume, c.MaxWeight, c.UsedVolume, c.UsedWeight)
	}
}

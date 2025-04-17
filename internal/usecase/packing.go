package usecase

import (
	"github.com/tamvanum/go-binpacking/internal/domain"
)

func GreedyPack(boxes []domain.Box, limit domain.Container) []domain.Container {
	var containers []domain.Container
	nextID := 1
	for _, box := range boxes {
		if !assingToFirstFit(containers, box) {
			containers = append(containers, newContainerWithBox(nextID, box, limit))
			nextID++
		}
	}
	return containers
}

func assingToFirstFit(containers []domain.Container, box domain.Box) bool {
	for i := range containers {
		c := &containers[i]
		if c.UsedVolume+box.Volume <= c.MaxWeight &&
			c.UsedVolume+box.Weight <= c.MaxVolume {
			c.Boxes = append(c.Boxes, box)
			c.UsedVolume += box.Volume
			c.UsedWeight += box.Weight
			return true
		}
	}
	return false
}

func newContainerWithBox(id int, box domain.Box, limit domain.Container) domain.Container {
	return domain.Container{
		ID:         id,
		Boxes:      []domain.Box{box},
		MaxWeight:  limit.MaxWeight,
		MaxVolume:  limit.MaxVolume,
		UsedWeight: box.Weight,
		UsedVolume: box.Volume,
	}
}

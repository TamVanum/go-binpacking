package repository

import "github.com/tamvanum/go-binpacking/internal/domain"

type BoxInMemoryRepository struct {
	boxes     []domain.Box
	idCounter int
}

func NewInMemoryBoxRepository() *BoxInMemoryRepository {
	return &BoxInMemoryRepository{boxes: []domain.Box{}}
}

func (r *BoxInMemoryRepository) Get() []domain.Box {
	return r.boxes
}

func (r *BoxInMemoryRepository) Add(box domain.Box) {
	r.idCounter++
	box.ID = r.idCounter
	r.boxes = append(r.boxes, box)
}

func (r *BoxInMemoryRepository) Clear() {
	r.boxes = nil
	r.idCounter = 0
}

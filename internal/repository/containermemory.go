package repository

import "github.com/tamvanum/go-binpacking/internal/domain"

type ContainerInMemoryRepository struct {
	containers []domain.Container
}

func NewContainerInMemoryRepository() *ContainerInMemoryRepository {
	return &ContainerInMemoryRepository{containers: []domain.Container{}}
}

func (r *ContainerInMemoryRepository) Get() []domain.Container {
	return r.containers
}

func (r *ContainerInMemoryRepository) Add(container domain.Container) {
	r.containers = append(r.containers, container)
}

func (r *ContainerInMemoryRepository) Clear() {
	r.containers = nil
}

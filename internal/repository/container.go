package repository

import "github.com/tamvanum/go-binpacking/internal/domain"

type ContainerRepository interface {
	Get() []domain.Container
	Add(container domain.Container)
	Clear()
}

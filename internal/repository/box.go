package repository

import "github.com/tamvanum/go-binpacking/internal/domain"

type BoxRepository interface {
	Get() []domain.Box
	Add(box domain.Box)
	Clear()
}

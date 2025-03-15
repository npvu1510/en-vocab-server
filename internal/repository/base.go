package repository

import (
	baseRepo "github.com/saoladigital/sld-go-common/repository"
)

type IBaseRepo[T baseRepo.Entity] interface {
	baseRepo.CURDBase[T]
}

type BaseRepo[T baseRepo.Entity] struct {
	baseRepo.CURDBase[T]
}

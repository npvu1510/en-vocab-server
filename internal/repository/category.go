package repository

import (
	"context"

	"github.com/npvu1510/en-vocab-server/internal/dto"
	"github.com/npvu1510/en-vocab-server/internal/model"
	baseRepo "github.com/saoladigital/sld-go-common/repository"
	"gorm.io/gorm"
)

type ICategoryRepo interface {
	GetMany(ctx context.Context, reqData dto.ListReqData) ([]model.Category, error)
}

type CategoryRepo struct {
	baseRepo.CURDBase[model.Category]
}

func NewCategoryRepo(db *gorm.DB) ICategoryRepo {
	return &CategoryRepo{CURDBase: baseRepo.NewCURDBaseImpl[model.Category](db)}

}

func (c *CategoryRepo) GetMany(ctx context.Context, reqData dto.ListReqData) ([]model.Category, error) {
	categories, err := c.CURDBase.Find(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

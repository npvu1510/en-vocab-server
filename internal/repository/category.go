package repository

import (
	"github.com/npvu1510/en-vocab-server/internal/dto"
	"github.com/npvu1510/en-vocab-server/internal/model"
	"gorm.io/gorm"
)

type ICategoryRepo interface {
	GetMany(reqData dto.ListReqData) ([]*model.Category, error)
}

type CategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) ICategoryRepo {
	return &CategoryRepo{db: db}
}

func (r *CategoryRepo) GetMany(reqData dto.ListReqData) ([]*model.Category, error) {
	var categories []*model.Category

	// PAGING
	if reqData.Page <= 0 {
		reqData.Page = 1
	}

	if reqData.Limit <= 0 {
		reqData.Limit = 10
	}
	offset := (reqData.Page - 1) * reqData.Limit

	err := r.db.Offset(int(offset)).Limit(int(reqData.Limit)).Find(&categories).Error

	return categories, err
}

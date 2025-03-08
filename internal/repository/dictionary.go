package repository

import (
	"github.com/npvu1510/en-vocab-server/internal/dto"
	"github.com/npvu1510/en-vocab-server/internal/model"
	"gorm.io/gorm"
)

type IDictionaryRepo interface {
	GetMany(reqData dto.ListReqData) ([]*model.Dictionary, error)
}

type DictionaryRepo struct {
	db *gorm.DB
}

func NewDictionaryRepo(db *gorm.DB) IDictionaryRepo {
	return &DictionaryRepo{db: db}
}

func (r *DictionaryRepo) GetMany(reqData dto.ListReqData) ([]*model.Dictionary, error) {
	var dictionaries []*model.Dictionary

	// ORDER
	var result = r.db.Order("definition asc")

	// PAGING
	if reqData.Page <= 0 {
		reqData.Page = 1
	}

	if reqData.Limit <= 0 {
		reqData.Limit = 10
	}
	offset := (reqData.Page - 1) * reqData.Limit
	result.Offset(int(offset)).Limit(int(reqData.Limit))

	err := result.Find(&dictionaries).Error

	return dictionaries, err
}

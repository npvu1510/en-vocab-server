package repository

import (
	"context"
	"math"

	"github.com/npvu1510/en-vocab-server/internal/dto"
	"github.com/npvu1510/en-vocab-server/internal/model"
	baseRepo "github.com/saoladigital/sld-go-common/repository"
	"gorm.io/gorm"
)

type IDictionaryRepo interface {
	IBaseRepo[model.Dictionary]
	GetMany(ctx context.Context, reqData dto.ListReqData) ([]model.Dictionary, int, error)
	// GetManyWithCategoryId(reqData dto.ListReqData) ([]*model.Dictionary, int, error)
}

type DictionaryRepo struct {
	baseRepo.CURDBase[model.Dictionary]
}

func NewDictionaryRepo(db *gorm.DB) IDictionaryRepo {
	curdBase := baseRepo.NewCURDBaseImpl[model.Dictionary](db)
	return &DictionaryRepo{CURDBase: curdBase}
}

func (d *DictionaryRepo) GetMany(ctx context.Context, reqData dto.ListReqData) ([]model.Dictionary, int, error) {
	// ################################# CONDITIONS #################################
	// ORDER
	var orderOpts baseRepo.QueryOption = func(db *gorm.DB) {
		db.Order("definition asc")
	}

	// PAGING
	if reqData.Page <= 0 {
		reqData.Page = 1
	}

	if reqData.Limit <= 0 {
		reqData.Limit = 10
	}

	var offsetOpts baseRepo.QueryOption = func(db *gorm.DB) {
		db.Offset(int((reqData.Page - 1) * reqData.Limit))
	}

	var limitOpts baseRepo.QueryOption = func(db *gorm.DB) {
		db.Limit(int(reqData.Limit))
	}

	// PRELOAD
	var preloadOpts baseRepo.QueryOption = func(db *gorm.DB) {
		db.Preload("Categories", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		})
	}

	// ################################# QUERY #################################
	totalItems, err := d.CURDBase.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	var totalPages = math.Ceil(float64(totalItems) / float64(reqData.Limit))

	dictionaries, err := d.CURDBase.Find(ctx, &preloadOpts, &orderOpts, &offsetOpts, &limitOpts)
	if err != nil {
		return nil, 0, err
	}
	return dictionaries, int(totalPages), err
}

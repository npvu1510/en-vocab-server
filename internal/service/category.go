package service

import (
	"context"

	"github.com/npvu1510/en-vocab-server/internal/dto"
	"github.com/npvu1510/en-vocab-server/internal/model"
	"github.com/npvu1510/en-vocab-server/internal/repository"
)

type ICategoryService interface {
	GetCategories(ctx context.Context, reqData dto.ListReqData) ([]model.Category, error)
}

type CategoryService struct {
	Repo repository.ICategoryRepo
}

func NewCategoryService(repo repository.ICategoryRepo) ICategoryService {
	return &CategoryService{Repo: repo}
}

func (s *CategoryService) GetCategories(ctx context.Context, reqData dto.ListReqData) ([]model.Category, error) {
	return s.Repo.GetMany(ctx, reqData)
}

package service

import (
	"context"

	"github.com/npvu1510/en-vocab-server/internal/dto"
	"github.com/npvu1510/en-vocab-server/internal/model"
	"github.com/npvu1510/en-vocab-server/internal/repository"
)

type IDictionaryService interface {
	GetDictionaries(ctx context.Context, reqData dto.ListReqData) ([]model.Dictionary, int, error)
}

type DictionaryService struct {
	Repo repository.IDictionaryRepo
}

func NewDictionaryService(repo repository.IDictionaryRepo) IDictionaryService {
	return &DictionaryService{Repo: repo}
}

func (s *DictionaryService) GetDictionaries(ctx context.Context, reqData dto.ListReqData) ([]model.Dictionary, int, error) {
	return s.Repo.GetMany(ctx, reqData)
}

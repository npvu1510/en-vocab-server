package service

import (
	"github.com/npvu1510/en-vocab-server/internal/dto"
	"github.com/npvu1510/en-vocab-server/internal/model"
	"github.com/npvu1510/en-vocab-server/internal/repository"
)

type IDictionaryService interface {
	GetDictionaries(reqData dto.ListReqData) ([]*model.Dictionary, error)
}

type DictionaryService struct {
	Repo repository.IDictionaryRepo
}

func NewDictionaryService(repo repository.IDictionaryRepo) IDictionaryService {
	return &DictionaryService{Repo: repo}
}

func (s *DictionaryService) GetDictionaries(reqData dto.ListReqData) ([]*model.Dictionary, error) {
	return s.Repo.GetMany(reqData)
}

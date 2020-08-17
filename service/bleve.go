package service

import (
	"bleve-indexing/internal/def"
	"bleve-indexing/internal/models"
	"bleve-indexing/internal/utils"
	"errors"

	"bleve-indexing/internal/bmapping"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
)

type Service struct {
	Kvstore       string
	IndexPath     string
	IndexType     string
	IndexMapping  mapping.IndexMapping
	FieldsMapping map[string]*mapping.FieldMapping
}

//RegisterFieldMapping registers fields mapping to the service
func (s *Service) RegisterFieldMapping() {
	s.FieldsMapping = def.TypeFieldMapping
}

//CreateIndex creates new index at specified path
func (s *Service) CreateIndex(name string) (bleve.Index, error) {

	err := utils.EnsureDir(s.IndexPath)
	if err != nil {
		return nil, err
	}

	index, err := bleve.NewUsing(s.IndexPath+"/"+name, s.IndexMapping, s.IndexType, s.Kvstore, nil)
	if err != nil {
		return nil, err
	}

	return index, nil
}

//BuildIndexMapping adds fields mapping and creates index mapping
func (s *Service) BuildIndexMapping(table models.Table) error {

	if len(s.FieldsMapping) == 0 {
		return errors.New("Unregistered fields mapping")
	}

	tableMapping, err := bmapping.FieldsMapping(table, s.FieldsMapping)
	if err != nil {
		return err
	}

	indexMapping := bleve.NewIndexMapping()
	indexMapping.AddDocumentMapping("table", tableMapping)
	indexMapping.DefaultAnalyzer = "en"

	s.IndexMapping = indexMapping

	return nil
}

package service

import (
	"bleve-indexing/internal/models"
	"bleve-indexing/internal/utils"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
)

type Service struct {
	Kvstore      string
	IndexPath    string
	IndexType    string
	IndexMapping mapping.IndexMapping
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
func (s *Service) BuildIndexMapping(data models.Table) error {

	//TODO Application specific Field Mapping

	tableMapping := bleve.NewDocumentMapping()

	indexMapping := bleve.NewIndexMapping()
	indexMapping.AddDocumentMapping("table", tableMapping)
	indexMapping.DefaultAnalyzer = "en"

	s.IndexMapping = indexMapping

	return nil
}

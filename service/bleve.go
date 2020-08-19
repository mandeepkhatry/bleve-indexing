package service

import (
	"bleve-indexing/internal/def"
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

//IndexRegister registers kvstore, indextype, indexpath and fields mapping to the service
func (s *Service) IndexRegister(kvStore string, indexType string, indexPath string) {
	s.Kvstore = kvStore
	s.IndexType = indexType
	s.IndexPath = indexPath
	s.FieldsMapping = def.TypeFieldMapping
}

//SearchRegister registers indexpath to the service
func (s *Service) SearchRegister(indexPath string) {
	s.IndexPath = indexPath
}

//Index creates new index at specified path and index the document.
func (s *Service) Index(name string, data map[string]interface{}) error {

	err := utils.EnsureDir(s.IndexPath)
	if err != nil {
		return err
	}

	index, err := bleve.NewUsing(s.IndexPath+"/"+name, s.IndexMapping, s.IndexType, s.Kvstore, nil)
	if err != nil {
		return err
	}

	index.Index(data["ID"].(string), data)

	return nil
}

//BuildIndexMapping adds fields mapping and creates index mapping
func (s *Service) BuildIndexMapping(data map[string]interface{}) error {

	if len(s.FieldsMapping) == 0 {
		return errors.New("Unregistered fields mapping")
	}

	tableMapping, err := bmapping.FieldsMapping(data, s.FieldsMapping)
	if err != nil {
		return err
	}

	indexMapping := bleve.NewIndexMapping()
	indexMapping.AddDocumentMapping("table", tableMapping)
	indexMapping.DefaultAnalyzer = "en"

	s.IndexMapping = indexMapping

	return nil
}

//Execute builds index mapping and index given document.
func (s *Service) Execute(name string, data map[string]interface{}) error {

	var err error

	err = s.BuildIndexMapping(data)
	if err != nil {
		return err
	}

	err = s.Index(name, data)
	if err != nil {
		return err
	}

	return nil
}

//TODO : Search  Current : Testing purpose only

func (s *Service) Search(name string, query string) (string, error) {

	var err error
	index, err := bleve.Open(s.IndexPath + "/" + name)
	if err != nil {
		return "", err
	}
	mquery := bleve.NewMatchQuery(query)
	searchRequest := bleve.NewSearchRequest(mquery)
	searchResult, err := index.Search(searchRequest)
	if err != nil {
		return "", err
	}
	return searchResult.Hits[0].ID, nil
}

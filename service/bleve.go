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

//CreateIndex returns index at specified path.
func (s *Service) CreateIndex(name string) (bleve.Index, error) {

	var index bleve.Index
	var err error

	path := s.IndexPath + "/" + name

	exists := utils.EnsureDir(path)

	if !exists {
		index, err = bleve.NewUsing(path, s.IndexMapping, s.IndexType, s.Kvstore, nil)
		if err != nil {
			return nil, err
		}
	} else {
		index, err = bleve.Open(path)
		if err != nil {
			return nil, err
		}
	}
	return index, err
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

//Index builds index mapping and index given document.
func (s *Service) Index(name string, data map[string]interface{}) error {

	var err error

	err = s.BuildIndexMapping(data)
	if err != nil {
		return err
	}

	index, err := s.CreateIndex(name)

	if err != nil {
		return err
	}

	err = index.Index(data["ID"].(string), data)
	index.Close()

	return err
}

//RunQuery executes query on the given store and returns set of matching ids.
func (s *Service) RunQuery(name string, query string, limit int, fields []string) ([]string, error) {

	var err error
	index, err := bleve.Open(s.IndexPath + "/" + name)
	if err != nil {
		return []string{}, err
	}
	mquery := bleve.NewQueryStringQuery(query)

	searchRequest := bleve.NewSearchRequest(mquery)
	searchRequest.Size = limit
	searchRequest.SortBy(fields)

	searchResult, err := index.Search(searchRequest)
	if err != nil {
		return []string{}, err
	}

	ids := make([]string, 0)

	for _, v := range searchResult.Hits {
		ids = append(ids, v.ID)
	}
	return ids, nil
}

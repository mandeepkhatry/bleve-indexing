package service

import (
	"bleve-indexing/internal/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
)

type Service struct {
	Kvstore      string
	IndexPath    string
	IndexType    string
	IndexMapping mapping.IndexMapping
}

//IndexRegister registers kvstore, indextype, indexpath and register indexing.
func (s *Service) IndexRegister(kvStore string, indexType string, indexPath string, name string, mappingPath string) error {

	s.Kvstore = kvStore
	s.IndexType = indexType
	s.IndexPath = indexPath

	path := s.IndexPath + "/" + name
	exists := utils.EnsureDir(path)

	if !exists {

		err := s.BuildIndexMapping(mappingPath)
		if err != nil {
			return err
		}

		_, err = s.CreateNewIndex(path)

		return err
	}

	//already registered
	return errors.New("index already registered")

}

//RegisterPath registers indexpath to the service
func (s *Service) RegisterPath(indexPath string) {
	s.IndexPath = indexPath
}

//CreateNewIndex creates index at specified path.
func (s *Service) CreateNewIndex(path string) (bleve.Index, error) {

	if s.IndexMapping == nil {
		return nil, errors.New("Unregistered index mapping")
	}

	index, err := bleve.NewUsing(path, s.IndexMapping, s.IndexType, s.Kvstore, nil)

	return index, err
}

//OpenIndex opens index at specified path.
func (s *Service) OpenIndex(path string) (bleve.Index, error) {

	index, err := bleve.Open(path)
	return index, err
}

//BuildIndexMapping builds index mapping according ot the given mapping.
func (s *Service) BuildIndexMapping(mappingPath string) error {

	mappingBytes, err := ioutil.ReadFile(mappingPath)
	if err != nil {
		return err
	}

	var indexMapping = mapping.NewIndexMapping()

	err = json.Unmarshal(mappingBytes, indexMapping)
	if err != nil {
		return err
	}

	s.IndexMapping = indexMapping

	return nil
}

//Index builds index mapping and index given document.
func (s *Service) Index(name string, data map[string]interface{}) error {

	var err error

	path := s.IndexPath + "/" + name

	fmt.Println(path)

	index, err := s.OpenIndex(path)

	if err != nil {
		return err
	}

	fmt.Println("THERE")
	err = index.Index(data["id"].(string), data)
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

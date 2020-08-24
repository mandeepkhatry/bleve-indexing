package service

import (
	"bleve-indexing/internal/def"
	"bleve-indexing/internal/imapping"
	"bleve-indexing/internal/utils"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
)

type IndexRegisterService struct {
	Kvstore      string
	IndexType    string
	Path         string
	IndexMapping mapping.IndexMapping
}

//NewIndexRegisterService returns new IndexService instance.
func NewIndexRegisterService(kvStore string, indexType string, indexPath string, name string) *IndexRegisterService {
	return &IndexRegisterService{
		Kvstore:   kvStore,
		IndexType: indexType,
		Path:      indexPath + "/" + name,
	}
}

//IndexRegister registers kvstore, indextype, indexpath and register indexing.
func (s *IndexRegisterService) IndexRegister(docmapping map[string]interface{}) error {

	if len(s.Kvstore) == 0 || len(s.IndexType) == 0 || len(s.Path) == 0 {
		return def.ErrIndexRegisterServiceUnRegistered
	}

	exists := utils.EnsureDir(s.Path)

	if !exists {

		err := s.BuildIndexMapping(docmapping)
		if err != nil {
			return err
		}

		err = s.CreateNewIndex(s.Path)

		return err
	}

	//already registered
	return def.ErrIndexRegistered

}

//CreateNewIndex creates index at specified path.
func (s *IndexRegisterService) CreateNewIndex(path string) error {

	if s.IndexMapping == nil {
		return def.ErrIndexMappingUnregistered
	}

	index, err := bleve.NewUsing(path, s.IndexMapping, s.IndexType, s.Kvstore, nil)
	index.Close()

	return err
}

//OpenIndex opens index at specified path.
func (s *IndexRegisterService) OpenIndex(path string) (bleve.Index, error) {

	index, err := bleve.Open(path)
	return index, err
}

//BuildIndexMapping builds index mapping according to the given mapping.
func (s *IndexRegisterService) BuildIndexMapping(docmapping map[string]interface{}) error {

	indexMapping, err := imapping.BuildMapping(docmapping)

	if err != nil {
		return err
	}

	//Validate indexmapping
	err = indexMapping.Validate()

	if err != nil {
		return err
	}

	s.IndexMapping = indexMapping

	return nil
}

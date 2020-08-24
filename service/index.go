package service

import (
	"bleve-indexing/internal/def"

	"github.com/blevesearch/bleve"
)

type IndexService struct {
	Path string
}

//NewIndexService returns new IndexService instance.
func NewIndexService(indexPath string, name string) *IndexService {
	return &IndexService{
		Path: indexPath + "/" + name,
	}
}

//Index opens index at given path and index given document.
func (s *IndexService) Index(data map[string]interface{}) error {

	var err error

	if len(s.Path) == 0 {
		return def.ErrIndexServiceUnRegistered
	}

	index, err := bleve.Open(s.Path)

	if err != nil {
		return err
	}

	err = index.Index(data["id"].(string), data)
	index.Close()

	return err
}

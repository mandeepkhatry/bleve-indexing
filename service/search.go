package service

import (
	"bleve-indexing/internal/def"

	"github.com/blevesearch/bleve"
)

type QueryService struct {
	Path  string
	Index bleve.Index
}

//NewQueryService returns new QueryService instance.
func NewQueryService(indexpath string, name string) (*QueryService, error) {
	var err error

	path := indexpath + "/" + name

	index, err := bleve.Open(path)
	if err != nil {
		return nil, err
	}

	return &QueryService{
		Path:  path,
		Index: index,
	}, nil
}

//RunQuery executes query on the given store and returns set of matching ids.
func (s *QueryService) RunQuery(iquery string, limit int, fields []string) ([]string, error) {

	if len(s.Path) == 0 || s.Index == nil {
		return []string{}, def.ErrQueryServiceUnRegistered
	}

	//Array of matched document ids
	ids := make([]string, 0)

	mquery := bleve.NewQueryStringQuery(iquery)

	searchRequest := bleve.NewSearchRequest(mquery)
	searchRequest.Size = limit
	searchRequest.SortBy(fields)

	searchResult, err := s.Index.Search(searchRequest)
	if err != nil {
		return []string{}, err
	}

	for _, v := range searchResult.Hits {
		ids = append(ids, v.ID)
	}

	return ids, nil

}

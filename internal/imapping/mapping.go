package imapping

import (
	"encoding/json"

	"github.com/blevesearch/bleve/mapping"
)

//BuildMapping builds index mapping on basis of given document mapping.
func BuildMapping(docmapping map[string]interface{}) (*mapping.IndexMappingImpl, error) {

	mappingBytes, _ := json.Marshal(docmapping)

	var indexMapping = mapping.NewIndexMapping()

	err := indexMapping.UnmarshalJSON(mappingBytes)

	return indexMapping, err
}

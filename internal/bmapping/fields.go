package bmapping

import (
	"bleve-indexing/internal/def"
	"bleve-indexing/internal/models"
	"encoding/json"
	"fmt"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
)

//FieldsMapping maps fields to documment mapping
func FieldsMapping(table models.Table, fmapping map[string]*mapping.FieldMapping) (*mapping.DocumentMapping, error) {
	var data map[string]interface{}
	err := json.Unmarshal(table.Multitenant_Data, &data)
	if err != nil {
		return nil, err
	}

	tableMapping := bleve.NewDocumentMapping()

	for key, value := range data {
		dataType := fmt.Sprintf("%T", value)
		appSpecificType := def.IndexFieldType[dataType]
		fieldMapping := fmapping[appSpecificType]
		tableMapping.AddFieldMappingsAt(key, fieldMapping)
	}

	return tableMapping, nil

}

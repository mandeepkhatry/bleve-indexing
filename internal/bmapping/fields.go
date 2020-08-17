package bmapping

import (
	"bleve-indexing/internal/def"
	"bleve-indexing/internal/models"
	"encoding/json"
	"fmt"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
)

func FieldsMapping(table models.Table) (*mapping.DocumentMapping, error) {
	var data map[string]interface{}
	err := json.Unmarshal(table.Multitenant_Data, &data)
	if err != nil {
		return nil, err
	}

	tableMapping := bleve.NewDocumentMapping()

	for key, value := range data {
		dataType := fmt.Sprintf("%T", value)
		appSpecificType := def.IndexFieldType[dataType]
		fieldMapping := def.TypeFieldMapping[appSpecificType]
		tableMapping.AddFieldMappingsAt(key, fieldMapping)
	}

	return tableMapping, nil

}

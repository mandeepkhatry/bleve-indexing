package bmapping

import (
	"bleve-indexing/internal/def"
	"bleve-indexing/internal/utils"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
)

//FieldsMapping maps fields to documment mapping
func FieldsMapping(data map[string]interface{}, fmapping map[string]*mapping.FieldMapping) (*mapping.DocumentMapping, error) {

	tableMapping := bleve.NewDocumentMapping()

	for key, value := range data {
		dataType := utils.FindType(value)
		appSpecificType := def.IndexFieldType[dataType]
		fieldMapping := fmapping[appSpecificType]
		tableMapping.AddFieldMappingsAt(key, fieldMapping)
	}

	return tableMapping, nil

}

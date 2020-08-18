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

	docMapping := bleve.NewDocumentMapping()

	for key, value := range data {
		dataType := utils.FindType(value)
		appSpecificType := def.IndexFieldType[dataType]

		if appSpecificType == "DOC" {
			var err error
			docMapping, err = FieldsMapping(value.(map[string]interface{}), fmapping)
			if err != nil {
				return nil, err
			}
			tableMapping.AddSubDocumentMapping("Data", docMapping)
		} else {
			fieldMapping := fmapping[appSpecificType]
			tableMapping.AddFieldMappingsAt(key, fieldMapping)
		}
	}

	return tableMapping, nil

}

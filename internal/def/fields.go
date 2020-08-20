package def

import (
	"github.com/blevesearch/bleve/mapping"

	"github.com/blevesearch/bleve"
)

//TypeFieldMapping returns type specific field mapping for given application specific type
var TypeFieldMapping = map[string]*mapping.FieldMapping{
	"TEXT":      GetTextFieldMapping(),
	"NUMERIC":   GetNumericFieldMapping(),
	"BOOLEAN":   GetBooleanFieldMapping(),
	"DATE_TIME": GetDateTimeFieldMapping(),
	"GEO_POINT": GetGeoFieldMapping(),
}

//GetTextFieldMapping returns Text Field Mapping
func GetTextFieldMapping() *mapping.FieldMapping {
	textFieldMapping := bleve.NewTextFieldMapping()
	textFieldMapping.Store = false
	return textFieldMapping
}

//GetNumericFieldMapping returns Numeric Field Mapping
func GetNumericFieldMapping() *mapping.FieldMapping {
	numericFieldMapping := bleve.NewNumericFieldMapping()
	numericFieldMapping.Store = false
	return numericFieldMapping
}

//GetBooleanFieldMapping returns Boolean Field Mapping
func GetBooleanFieldMapping() *mapping.FieldMapping {
	boolFieldMapping := bleve.NewBooleanFieldMapping()
	boolFieldMapping.Store = false
	return boolFieldMapping
}

//GetDateTimeFieldMapping returns Date Time Field Mapping
func GetDateTimeFieldMapping() *mapping.FieldMapping {
	datetimeFieldMapping := bleve.NewDateTimeFieldMapping()
	datetimeFieldMapping.Store = false
	return datetimeFieldMapping
}

//GetGeoFieldMapping returns Geo Field Mapping
func GetGeoFieldMapping() *mapping.FieldMapping {
	geoFieldMapping := bleve.NewGeoPointFieldMapping()
	geoFieldMapping.Store = false
	return geoFieldMapping
}

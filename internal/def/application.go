package def

//IndexFieldType represents application specific data type.
var IndexFieldType = map[string]string{
	"string":    "TEXT",
	"int":       "NUMERIC",
	"int32":     "NUMERIC",
	"int64":     "NUMERIC",
	"float32":   "NUMERIC",
	"float64":   "NUMERIC",
	"bool":      "BOOLEAN",
	"time.Time": "DATE_TIME",
	"geo_point": "GEO_POINT",
}

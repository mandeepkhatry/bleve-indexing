# Bleve-Indexing
Full-text Search and Indexing.

## Building Mapping

```
{
    "types": {
        "example": {
            "enabled": true,
            "dynamic": true,
            "properties": {
                "name": {
                    "enabled": true,
                    "dynamic": true,
                    "fields": [{
                        "type": "text",
                        "analyzer": "en",
                        "store": true,
                        "index": true,
                        "include_term_vectors": true,
                        "include_in_all": true
                    }],
                    "default_analyzer": "en"
                },
                "timestamp": {
                    "enabled": true,
                    "dynamic": true,
                    "fields": [{
                        "type": "datetime",
                        "store": true,
                        "index": true,
                        "include_in_all": true
                    }],
                    "default_analyzer": ""
                }
            },
            "default_analyzer": "en"
        }
    },
    "default_mapping": {
        "enabled": true,
        "dynamic": true,
        "default_analyzer": "standard"
    },
    "type_field": "_type",
    "default_type": "_default",
    "default_analyzer": "standard",
    "default_datetime_parser": "dateTimeOptional",
    "default_field": "_all",
    "store_dynamic": true,
    "index_dynamic": true,
    "analysis": {
        "analyzers": {},
        "char_filters": {},
        "tokenizers": {},
        "token_filters": {},
        "token_maps": {}
    }
}
```

## Index Register

```
s := service.Service{}

//Register Kvstore, index type, path to store and registers mappings.
err = s.IndexRegister("scorch", "scorch", "store", "nica.employee.1", docmapping)

if err != nil {
    panic(err)
}
```


## Indexing

```

indexService := service.Service{}

//Register Path
indexService.RegisterPath("store")

//Index document
err = indexService.Index("nica.employee.1", document)
if err != nil {
    panic(err)
}

fmt.Println("Indexing Finished.", document["id"])

```

## Executing Query

```
searchService := service.Service{}
searchService.RegisterPath("store")

limit := 100
sortFields := []string{"id}
query := "created_by:admin"

//Run Query
id, err := searchService.RunQuery("nica.employee.1", query, limit, sortFields)
if err != nil {
    panic(err)
}

fmt.Println("Document IDs : ", id)

```
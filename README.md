# Bleve-Indexing
Full-text Search and Indexing.

## Building Mapping

```
{
    "types": {
        "document": {
            "enabled": true,
            "dynamic": true,
            "properties": {
                "id": {
                    "enabled": true,
                    "dynamic": true,
                    "fields": [{
                        "type": "text",
                        "analyzer": "en",
                        "store": false,
                        "index": true,
                        "include_term_vectors": true,
                        "include_in_all": true
                    }],
                    "default_analyzer": ""
                },
                "group_label": {
                    "enabled": true,
                    "dynamic": true,
                    "fields": [{
                        "type": "text",
                        "analyzer": "en",
                        "store": false,
                        "index": true,
                        "include_term_vectors": true,
                        "include_in_all": true
                    }],
                    "default_analyzer": ""

                },
                "status": {
                    "enabled": true,
                    "dynamic": true,
                    "fields": [{
                        "type": "number",
                        "store": false,
                        "index": true,
                        "include_in_all": true
                    }],
                    "default_analyzer": ""

                },
                "created_at": {
                    "enabled": true,
                    "dynamic": true,
                    "fields": [{
                        "type": "datetime",
                        "store": false,
                        "index": true,
                        "include_in_all": true
                    }],
                    "default_analyzer": ""
                },
                "created_by": {
                    "enabled": true,
                    "dynamic": true,
                    "fields": [{
                        "type": "text",
                        "analyzer": "en",
                        "store": false,
                        "index": true,
                        "include_term_vectors": true,
                        "include_in_all": true
                    }],
                    "default_analyzer": ""
                },
                "last_modified_at": {
                    "enabled": true,
                    "dynamic": true,
                    "fields": [{
                        "type": "datetime",
                        "store": false,
                        "index": true,
                        "include_in_all": true
                    }],
                    "default_analyzer": ""
                },
                "last_modified_by": {
                    "enabled": true,
                    "dynamic": true,
                    "fields": [{
                        "type": "text",
                        "analyzer": "en",
                        "store": false,
                        "index": true,
                        "include_term_vectors": true,
                        "include_in_all": true
                    }],
                    "default_analyzer": ""
                },
                "permission": {
                    "enabled": true,
                    "dynamic": true,
                    "fields": [{
                        "type": "number",
                        "store": false,
                        "index": true,
                        "include_in_all": true
                    }],
                    "default_analyzer": ""
                },

                "data": {
                    "enabled": true,
                    "dynamic": true,
                    "properties": {
                        "name": {
                            "enabled": true,
                            "dynamic": true,
                            "fields": [{
                                "type": "text",
                                "analyzer": "en",
                                "store": false,
                                "index": true,
                                "include_term_vectors": true,
                                "include_in_all": true
                            }],
                            "default_analyzer": "en"
                        },
                        "age": {
                            "enabled": true,
                            "dynamic": true,
                            "fields": [{
                                "type": "number",
                                "store": false,
                                "index": true,
                                "include_in_all": true
                            }],
                            "default_analyzer": ""
                        },
                        "work": {
                            "enabled": true,
                            "dynamic": true,
                            "fields": [{
                                "type": "text",
                                "analyzer": "en",
                                "store": false,
                                "index": true,
                                "include_term_vectors": true,
                                "include_in_all": true
                            }],
                            "default_analyzer": "en"
                        }
                    },
                    "default_analyzer": ""

                },
                "default_mapping": {
                    "enabled": true,
                    "dynamic": true,
                    "default_analyzer": ""
                }

            },
            "default_analyzer": ""
        }
    },
    "default_mapping": {
        "enabled": true,
        "dynamic": true,
        "default_analyzer": ""
    },
    "type_field": "type",
    "default_type": "document",
    "default_analyzer": "standard",
    "default_datetime_parser": "dateTimeOptional",
    "default_field": "_all",
    "store_dynamic": true,
    "index_dynamic": true,
    "analysis": {}
}
```

## Index Register

```

iRegisterService := service.NewIndexRegisterService("scorch", "scorch", "store", "nica.employee.1")
err = iRegisterService.IndexRegister(docmapping)
if err != nil {
    log.Fatal(err)
}
```


## Indexing

```
indexService := service.NewIndexService("store", "nica.employee.1")
err = indexService.Index(document)
if err != nil {
    log.Fatal(err)
}

```

## Executing Query

```
queryService, err := service.NewQueryService("store", "nica.employee.1")
if err != nil {
    log.Fatal(err)
}

limit := 100
sortFields := []string{"id"}
ids, err := queryService.RunQuery("data.name:Mandeep", limit, sortFields)
if err != nil {
    log.Fatal(err)
}
fmt.Println("Document IDs : ", ids)

```
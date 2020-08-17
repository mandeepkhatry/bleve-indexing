package main

import (
	"bleve-indexing/internal/models"
	"bleve-indexing/service"
	"encoding/json"
	"fmt"
	"time"
)

func main() {

	multitenantData := map[string]interface{}{
		"name": "Mandeep Khatry",
		"age":  22,
		"work": "developer",
	}

	byteData, err := json.Marshal(multitenantData)
	if err != nil {
		panic(err)
	}

	data := models.Table{
		ID:               "123",
		Namespace:        "NICA",
		Collection:       "Employee",
		Partition:        0,
		Data_Schema:      2,
		Created_At:       time.Now(),
		Created_By:       "admin",
		Multitenant_Data: byteData,
	}

	fmt.Println(data)

	//test
	s := service.Service{Kvstore: "scorch", IndexType: "scorch", IndexPath: "store"}
	s.RegisterFieldMapping()

	fmt.Println(s.FieldsMapping)
}

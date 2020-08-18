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

	data := models.DynamicTable{
		ID:               "123",
		Namespace:        "NICA",
		Collection:       "Employee",
		Group_Label:      "A",
		Partition:        0,
		Status:           0,
		Stage:            0,
		Created_At:       time.Now(),
		Created_By:       "admin",
		Last_Modified_At: time.Now(),
		Last_Modified_By: "admin",
		Times_Modified:   2,
		Permission:       1,
		Data_Schema:      2,
		Multitenant_Data: byteData,
	}

	fmt.Println(data)

	var mData map[string]interface{}
	err = json.Unmarshal(data.Multitenant_Data, &mData)
	if err != nil {
		panic(err)
	}

	document := map[string]interface{}{
		"ID":               data.ID,
		"Group_Label":      data.Group_Label,
		"Status":           data.Status,
		"Stage":            data.Stage,
		"Created_At":       data.Created_At,
		"Created_By":       data.Created_By,
		"Last_Modified_At": data.Last_Modified_At,
		"Last_Modified_By": data.Last_Modified_By,
		"Permission":       data.Permission,
		"Data":             mData,
	}

	fmt.Println(document["ID"])

	//test
	s := service.Service{Kvstore: "scorch", IndexType: "scorch", IndexPath: "store"}
	s.RegisterFieldMapping()

	fmt.Println(s.FieldsMapping)
}

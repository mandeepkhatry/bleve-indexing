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

	s := service.Service{}
	//Register Kvstore, index type, store path and registers field mappings.
	s.IndexRegister("scorch", "scorch", "store")
	//Execute Indexing Service
	s.Execute("nica.employee.1", document)

	fmt.Println("Indexing Finished.")

	// //Testing Search
	// s := service.Service{}
	// s.SearchRegister("store")
	// id, err := s.Search("nica.employee.1", "ID=123")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Document ID : ", id)

}

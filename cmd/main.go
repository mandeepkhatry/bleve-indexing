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
		"name":       "Mandeep Khatry",
		"age":        22,
		"work":       "developer",
		"isEmployee": false,
	}

	byteData, err := json.Marshal(multitenantData)
	if err != nil {
		panic(err)
	}

	data := models.DynamicTable{
		ID:               "1",
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

	// //Execute Indexing Service
	// s := service.Service{}
	// //Register Kvstore, index type, store path and registers field mappings.
	// s.IndexRegister("scorch", "scorch", "store")

	// document := map[string]interface{}{
	// 	"ID":               data.ID,
	// 	"Group_Label":      data.Group_Label,
	// 	"Status":           data.Status,
	// 	"Stage":            data.Stage,
	// 	"Created_At":       data.Created_At,
	// 	"Created_By":       data.Created_By,
	// 	"Last_Modified_At": data.Last_Modified_At,
	// 	"Last_Modified_By": data.Last_Modified_By,
	// 	"Permission":       data.Permission,
	// 	"Data":             mData,
	// }
	// err = s.Index("nica.employee.1", document)
	// if err != nil {
	// 	log.Println("Error : ", err)
	// 	os.Exit(1)
	// }
	// fmt.Println("Indexing Finished.", document["ID"])

	//Testing Search
	s := service.Service{}
	s.SearchRegister("store")

	//Execute a query in specified store with specific limit.
	id, err := s.RunQuery("nica.employee.1", "Data.age:>=22", 100, []string{"name"})
	if err != nil {
		panic(err)
	}
	fmt.Println("Document IDs : ", id)

}

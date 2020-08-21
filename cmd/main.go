package main

import (
	"bleve-indexing/internal/models"
	"bleve-indexing/service"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
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

	// document := map[string]interface{}{
	// 	"id":               data.ID,
	// 	"group_label":      data.Group_Label,
	// 	"status":           data.Status,
	// 	"stage":            data.Stage,
	// 	"created_at":       data.Created_At,
	// 	"created_by":       data.Created_By,
	// 	"last_modified_at": data.Last_Modified_At,
	// 	"last_modified_by": data.Last_Modified_By,
	// 	"permission":       data.Permission,
	// 	"data":             mData,
	// }

	//*******Testing*******
	mappingBytes, err := ioutil.ReadFile("mapping/test.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	var docmapping map[string]interface{}

	json.Unmarshal(mappingBytes, &docmapping)

	//Execute Indexing Service
	s := service.Service{}
	//Register Kvstore, index type, store path and registers field mappings.
	err = s.IndexRegister("scorch", "scorch", "store", "nica.employee.1", docmapping)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// indexService := service.Service{}
	// indexService.RegisterPath("store")
	// err = indexService.Index("nica.employee.1", document)
	// if err != nil {
	// 	log.Println("Error : ", err)
	// 	os.Exit(1)
	// }
	// fmt.Println("Indexing Finished.", document["id"])

	// //Testing Search
	// s := service.Service{}
	// s.RegisterPath("store")

	// //Execute a query in specified store with specific limit.

	// id, err := s.RunQuery("nica.employee.1", "id:1", 100, []string{"group_label"})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Document IDs : ", id)

}

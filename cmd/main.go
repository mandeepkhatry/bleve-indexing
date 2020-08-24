package main

import (
	"bleve-indexing/internal/models"
	"bleve-indexing/service"
	"encoding/json"
	"fmt"
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
		Group_Label:      "I am going to school",
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

	// //*******Testing*******
	mappingBytes, err := ioutil.ReadFile("mapping/nica-employee.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	var docmapping map[string]interface{}

	json.Unmarshal(mappingBytes, &docmapping)

	//Index Register
	iRegisterService := service.NewIndexRegisterService("scorch", "scorch", "store", "nica.employee.1")
	err = iRegisterService.IndexRegister(docmapping)
	if err != nil {
		log.Fatal(err)
	}

	//Indexing Document

	// ******Testing Indexing**************

	document := map[string]interface{}{
		"id":               data.ID,
		"group_label":      data.Group_Label,
		"status":           data.Status,
		"stage":            data.Stage,
		"created_at":       data.Created_At,
		"created_by":       data.Created_By,
		"last_modified_at": data.Last_Modified_At,
		"last_modified_by": data.Last_Modified_By,
		"permission":       data.Permission,
		"data":             mData,
	}

	indexService := service.NewIndexService("store", "nica.employee.1")
	err = indexService.Index(document)
	if err != nil {
		log.Fatal(err)
	}

	//Quering
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

}

package models

import "time"

//Table represents struct for Data Model for indexing.
type Table struct {
	ID               string
	Namespace        string
	Collection       string
	Group_Label      string
	Partition        int64
	Status           int64
	Stage            int64
	Created_At       time.Time
	Created_By       string
	Last_Modified_At time.Time
	Last_Modified_By string
	Times_Modified   int64
	Permission       int64
	Data_Schema      int64
	Multitenant_Data []byte
}

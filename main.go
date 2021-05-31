package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	db, err := NewDatabase()
	if err != nil {panic(err)}

	e := Example{
		Id:   "1",
		Name: "test",
		Age:  2,
	}
	b, err := json.Marshal(&e)
	if err != nil {panic(err)}
	err = db.Set(e.Id, b)
	if err != nil {panic(err)}

	b2, err := db.Get("1")
	if err != nil {panic(err)}
	var e2 Example
	json.Unmarshal(b2, &e2)

	fmt.Println(e2)
}

type Example struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Age int64 `json:"age"`
}

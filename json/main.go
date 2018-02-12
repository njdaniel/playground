package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type Account struct {
	Server	string 	`json:"server"`
}

// JSON types:
// numbers, boolean, strings
// bool for JSON boolean
// float64 for JSON numbers
// string for JSON string
// []interface{} for JSON array
// map[string]interface{} for JSON objects
// nil for JSON null
func main() {
	accounts := []byte(`{"server": "qa"}`)
	var f interface{}
	err := json.Unmarshal(accounts, &f)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%v %T\n", f, f)
	var account Account
	m := f.(map[string]string)
	//m := f.(map[string]interface{})
	account.Server = m["server"]
}

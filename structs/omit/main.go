package main

import (
	"encoding/json"
	"fmt"
)

// Person might not have middle names
type Person struct {
	First  string `json:"First"`
	Middle string `json:"Middle"`
	Age    int    `json:"Age"`
}

func main() {
	var p1 Person
	bs := []byte(`{"First": "Nick"}`)

	json.Unmarshal(bs, &p1)

	fmt.Println("----------------")
	fmt.Println(p1)
	fmt.Println(p1.First)
	fmt.Println(p1.Middle)
	fmt.Println(p1.Age)
}

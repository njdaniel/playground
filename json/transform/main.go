package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

type idata struct {
	One   string
	Two   string
	Three string
}

type odata struct {
	One string
	Two string
}

func main() {
	json := `{
		"k1" : "v1",
		"k2" : "v2"
		`

	value := gjson.Get(json, "k1")
	fmt.Println(value)

}

package main

import (
	"fmt"
	"encoding/json"
)

type King struct {
	Source string	`json:"source"`
	Dest []string	`json:"dest"`
}

func main() {
	jsonStr := "{\"source\": \"test\", \"dest\": [\"a\", \"b\"]}"
	// json, _ := simplejson.NewJson([]byte(jsonStr)).Decode(&King{})
	var k King
	json.Unmarshal([]byte(jsonStr), &k)
	fmt.Println(k)
}

func init() {
	fmt.Println("init func in main package")
}

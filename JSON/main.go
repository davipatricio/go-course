package main

import (
	"encoding/json"
	"fmt"
)

// Create a structure then marshal it to JSON
type person struct {
	Name         string `json:"name"`
	Age          int    `json:"age"`
	FavoriteFood string `json:"favoriteFood"`
}

func transformStringToJson() {
	var dat map[string]interface{}
	jsonStr := []byte(`{"Name":"Alfredo","Age":30,"FavoriteFood":"Pizza"}`)

	// Unmarshal the JSON to the map
	json.Unmarshal(jsonStr, &dat)

	fmt.Println(dat["Name"].(string))
	fmt.Println(dat)
}

func main() {
	p1 := person{"Alfredo", 30, "Pizza"}

	// Marshal the structure to JSON
	json, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	// Print the JSON
	fmt.Println(string(json))
	transformStringToJson()
}

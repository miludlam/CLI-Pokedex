package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMap(c *config) error {
	callAPI(c)
	return nil
}

func callAPI(c *config) {
	res, err := http.Get(c.Next)

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	var areaMap map[string]interface{}
	err = json.Unmarshal(body, &areaMap)
	if err != nil {
		log.Fatal(err)
	}

	// Update our config fields
	c.Previous = c.Next
	if areaMap["next"] != nil {
		c.Next = areaMap["next"].(string)
	} else {
		c.Next = ""
		fmt.Println("You're on the last page")
	}

	for _, area := range areaMap["results"].([]interface{}) {
		area := area.(map[string]interface{})
		fmt.Println(area["name"].(string))
	}
}

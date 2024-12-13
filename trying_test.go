package main

import (
	"testing"
)

func sampleJsonOutputParsingFunc() ([]map[string]interface{}) {

	jsonData := []map[string]interface{}{}
	temp := dailyScorePrioQueue.front
	for temp != nil {
		jsonData = append(jsonData, map[string]interface{}{
			"name":  temp.item.username,
			"score": temp.item.score,
		},
		)
		temp = temp.next
	}

	return jsonData
}

//if i rmb correctly this is an array of json objects, i am testing if name and score fields are populated correctly
func TestMapSomething(t *testing.T) {
	for _, mapEntry := range sampleJsonOutputParsingFunc() {
		if mapEntry["name"] == "" {
			t.Error("name field is not populated in the json object")
		}
		if mapEntry["score"] == "" {
			t.Error("score field is not populated in the json object")
		}
	}
}

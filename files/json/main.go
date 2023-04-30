package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ghodss/yaml"
)

type journey struct {
	Name     string `json:"name" yaml:"name"`
	Status   string `json:"status" yaml:"status"`
	Progress int    `json:"progress" yaml:"progress"`
}

type skill struct {
	Name     string    `json:"name" yaml:"name"`
	Level    int       `json:"level" yaml:"level"`
	Topic    string    `json:"topic" yaml:"topic"`
	Subject  string    `json:"subject" yaml:"subject"`
	Priority int       `json:"priority" yaml:"priority"`
	Journeys []journey `json:"journeys" yaml:"journeys"`
}

func convert_from_json_to_struct(some_json []byte) (skill, error) {
	var some_struct skill

	err := json.Unmarshal(some_json, &some_struct)
	if err != nil {
		log.Panicf("error while convert from json to struct: %s", err)
	}
	return some_struct, err
}

func convert_from_struct_to_json(some_struct skill) ([]byte, error) {
	some_json, err := json.MarshalIndent(&some_struct, "", "  ")
	if err != nil {
		log.Fatalf("error while convert from struct to json: %s", err)
	}
	return some_json, err
}

func convert_from_json_to_yaml(some_json []byte) ([]byte, error) {
	some_yaml, err := yaml.JSONToYAML(some_json)
	if err != nil {
		log.Fatalf("error while convert from struct to json: %s", err)
	}
	return some_yaml, err
}

// func check_if_key_exists(json_content []byte, target_key string) bool {
// 	return true
// }

func main() {
	// ---------- CONVERT[JSON:STRUCT] ----------
	var some_json = []byte(`{"name": "golang", "level": 1, "topic": "programming", "subject": "language", "priority": 1, "journeys": [{"name": "working-with-json", "status": "in-progress", "progress": 1}]}`)
	converted_json_to_struct, err := convert_from_json_to_struct(some_json)
	if err != nil {
		log.Fatalf("error while conversion: %s", err)
	}
	fmt.Printf("[from JSON to STRUCT]\n%+v\n\n", converted_json_to_struct)

	// ---------- CONVERT[STRUCT:JSON] ----------
	some_struct := skill{
		Name:     "golang",
		Level:    1,
		Topic:    "development",
		Subject:  "language",
		Priority: 1,
		Journeys: []journey{
			{
				Name:     "working-with-json",
				Status:   "in-progress",
				Progress: 1,
			},
		},
	}

	converted_struct_to_json, err := convert_from_struct_to_json(some_struct)
	if err != nil {
		log.Fatalf("error while conversion: %s", err)
	}
	fmt.Printf("[from STRUCT to JSON]\n%s\n", converted_struct_to_json)

	// ---------- CONVERT[JSON:YAML] ----------
	some_yaml, err := convert_from_json_to_yaml(converted_struct_to_json)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n[from JSON to YAML]\n%+s\n\n", string(some_yaml))

	// ---------- CHECK_IF_KEY_EXISTS[JSON] ----------
	// target_key := "Level"
	// check_result := check_if_key_exists(converted_struct_to_json, target_key)
	// fmt.Printf("\n[check if key exists]\n%s: %t\n\n", target_key, check_result)

	// ---------- GET_JSON_KEY_VALUE[JSON] ----------
}

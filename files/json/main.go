package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/ghodss/yaml"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

func check_if_key_exists(some_key string, some_json []byte) bool {
	json_content, _ := convert_from_json_to_struct(some_json)
	struct_values := reflect.ValueOf(json_content)
	var check_status bool

	capitalized_key := cases.Title(language.Und).String(some_key)
	if struct_values.FieldByName(capitalized_key).IsValid() {
		check_status = true
	} else {
		check_status = false
	}
	return check_status
}

func get_key_value(some_key string, some_json []byte) any {
	json_content, _ := convert_from_json_to_struct(some_json)
	capitalized_key := cases.Title(language.Und).String(some_key)
	struct_info := reflect.ValueOf(json_content)
	key_value := struct_info.FieldByName(capitalized_key)
	return key_value
}

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
	fmt.Printf("\n[from JSON to YAML]\n%+s", string(some_yaml))

	// ---------- CHECK_KEY_EXISTS ----------
	key_name := "level"
	key_exists := check_if_key_exists(key_name, converted_struct_to_json)
	fmt.Printf("\n[check if a key exists]\n%s: %t\n", key_name, key_exists)

	// ---------- GET_KEY_VALUE ----------
	some_key_name := "topic"
	key_value := get_key_value(some_key_name, converted_struct_to_json)
	fmt.Printf("\n[get a key value]\n%s: ", some_key_name)
	fmt.Println(key_value)
}

package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/ghodss/yaml"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type skill struct {
	Name     string `json:"name" yaml:"name"`
	Level    int    `json:"level" yaml:"level"`
	Topic    string `json:"topic" yaml:"topic"`
	Subject  string `json:"subject" yaml:"subject"`
	Priority int    `json:"priority" yaml:"priority"`
}

func convert_from_yaml_to_struct(some_yaml []byte) skill {
	var some_struct skill

	err := yaml.Unmarshal(some_yaml, &some_struct)
	if err != nil {
		panic(err)
	}
	return some_struct
}

func convert_from_struct_to_yaml(some_struct skill) []byte {
	some_yaml, err := yaml.Marshal(&some_struct)
	if err != nil {
		log.Fatalf("error while convert from struct to json: %s", err)
	}
	return some_yaml
}

func convert_from_yaml_to_json(some_yaml []byte) []byte {
	some_json, err := yaml.YAMLToJSON(some_yaml)
	if err != nil {
		log.Fatalf("error while convert from struct to json: %s", err)
	}
	return some_json
}

func check_if_key_exists(some_key string, some_yaml []byte) bool {
	yaml_content := convert_from_yaml_to_struct(some_yaml)
	struct_values := reflect.ValueOf(yaml_content)
	var check_status bool

	capitalized_key := cases.Title(language.Und).String(some_key)
	if struct_values.FieldByName(capitalized_key).IsValid() {
		check_status = true
	} else {
		check_status = false
	}
	return check_status
}

func get_key_value(some_key string, some_yaml []byte) any {
	yaml_content := convert_from_yaml_to_struct(some_yaml)
	capitalized_key := cases.Title(language.Und).String(some_key)
	struct_info := reflect.ValueOf(yaml_content)
	key_value := struct_info.FieldByName(capitalized_key)
	return key_value
}

func main() {
	yaml_content := "name: golang\nlevel: 1\ntopic: programming\nsubject: language\npriority: 1"
	var some_yaml = []byte(yaml_content)

	// ---------- CONVERT[YAML:STRUCT] ----------
	converted_yaml_to_struct := convert_from_yaml_to_struct(some_yaml)
	fmt.Printf("[from YAML to STRUCT]\n%+v\n", converted_yaml_to_struct)

	// ---------- CONVERT[STRUCT:YAML] ----------
	converted_struct_struct_to_yaml := convert_from_struct_to_yaml(converted_yaml_to_struct)
	fmt.Printf("\n[from STRUCT to YAML]\n%s", converted_struct_struct_to_yaml)

	// ---------- CONVERT[YAML:JSON] ----------
	converted_yaml_to_json := convert_from_yaml_to_json(converted_struct_struct_to_yaml)
	fmt.Printf("\n[from YAML to JSON]\n%s", converted_yaml_to_json)

	// ---------- CHECK_KEY_EXISTS ----------
	key_name := "level"
	key_exists := check_if_key_exists(key_name, converted_struct_struct_to_yaml)
	fmt.Printf("\n\n[check if a key exists]\n%s: %t", key_name, key_exists)

	// ---------- GET_KEY_VALUE ----------
	some_key_name := "topic"
	key_value := get_key_value(some_key_name, converted_struct_struct_to_yaml)
	fmt.Printf("\n\n[get a key value]\n%s: ", some_key_name)
	fmt.Println(key_value)

}

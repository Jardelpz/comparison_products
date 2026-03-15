package models

import (
	"reflect"
	"strings"
)

type Product struct {
	ID          string         `json:"id"`
	Category    string         `json:"category"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Size        string         `json:"size"`
	Weight      float64        `json:"weight"`
	Color       string         `json:"color"`
	Specs       map[string]any `json:"specs"`
}

func GetProductDefaultFields() []string {
	t := reflect.TypeOf(Product{})

	keys := make([]string, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}

		jsonTag := f.Tag.Get("json")
		if idx := strings.IndexByte(jsonTag, ','); idx >= 0 {
			jsonTag = jsonTag[:idx]
		}

		if jsonTag == "id" || jsonTag == "specs" {
			continue
		}
		keys = append(keys, jsonTag)
	}

	return keys
}

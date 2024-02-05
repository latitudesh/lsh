package utils

import (
	"reflect"
	"strings"
)

func AssignValue(ptr interface{}, jsonName string, value interface{}) {
	// Get the reflect.Value of the struct using the pointer
	objValue := reflect.ValueOf(ptr).Elem()

	// Loop through the fields of the struct
	for i := 0; i < objValue.NumField(); i++ {
		field := objValue.Type().Field(i)

		// Check if the JSON tag matches the specified JSON field name
		tag := field.Tag.Get("json")
		if strings.Split(tag, ",")[0] == jsonName {
			// Get the reflect.Value of the field
			fieldValue := objValue.FieldByName(field.Name)

			// Check if the field is valid and can be set
			if fieldValue.IsValid() && fieldValue.CanSet() {
				// Convert the value to the appropriate type
				valueToSet := reflect.ValueOf(value)

				// Set the value of the field
				fieldValue.Set(valueToSet)
			}
			return
		}
	}
}

func GetFieldValue(ptr interface{}, jsonName string) reflect.Value {
	// Get the reflect.Value of the struct using the pointer
	objValue := reflect.ValueOf(ptr).Elem()

	// Loop through the fields of the struct
	for i := 0; i < objValue.NumField(); i++ {
		field := objValue.Type().Field(i)

		// Check if the JSON tag matches the specified JSON field name
		tag := field.Tag.Get("json")
		if strings.Split(tag, ",")[0] == jsonName {
			// Get the reflect.Value of the field
			return objValue.FieldByName(field.Name)
		}
	}

	return objValue
}

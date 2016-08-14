package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// This example shows the basic usage of the package: Create an encoder,
// transmit some values, receive them with a decoder.
func main() {
	csv := map[string]string{"Name": "ethan", "Age": "30", "Test": "true"}

	type man struct {
		Name string
		Age  int
		Test bool
	}

	var m man
	decode(&m, csv)

	fmt.Println(m)
}

func decode(result interface{}, csv map[string]string) {
	val := reflect.ValueOf(result)
	if val.Kind() != reflect.Ptr {
		panic("result must be a pointer")
	}

	val = val.Elem()
	if !val.CanAddr() {
		panic("result must be addressable (a pointer)")
	}

	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldName := field.Name
		fieldKind := field.Type.Kind()
		valueField := val.Field(i)

		if csvValue, ok := csv[fieldName]; ok {
			switch {
			case fieldKind == reflect.String:
				valueField.SetString(csvValue)
			case fieldKind == reflect.Bool:
				if b, err := strconv.ParseBool(csvValue); err == nil {
					valueField.SetBool(b)
				} else {
					panic(fmt.Sprintf("The value %s could NOT be converted to bool", csvValue))
				}
			case fieldKind == reflect.Int:
				if i, err := strconv.ParseInt(csvValue, 10, 32); err == nil {
					valueField.SetInt(i)
				} else {
					panic(fmt.Sprintf("The value %s could NOT be converted to int64", csvValue))
				}
			case fieldKind == reflect.Uint:
				if u, err := strconv.ParseUint(csvValue, 10, 32); err == nil {
					valueField.SetUint(u)
				} else {
					panic(fmt.Sprintf("The value %s could NOT be converted to uint64", csvValue))
				}
			case fieldKind == reflect.Float64:
				if f, err := strconv.ParseFloat(csvValue, 32); err == nil {
					valueField.SetFloat(f)
				} else {
					panic(fmt.Sprintf("The value %s could NOT be converted to float64", csvValue))
				}
			}
		} else {
			panic(fmt.Sprintf("There is NO value of column %s in the csv file", fieldName))
		}
	}
}

func getKind(val reflect.Value) reflect.Kind {
	kind := val.Kind()

	switch {
	case kind >= reflect.Int && kind <= reflect.Int64:
		return reflect.Int
	case kind >= reflect.Uint && kind <= reflect.Uint64:
		return reflect.Uint
	case kind >= reflect.Float32 && kind <= reflect.Float64:
		return reflect.Float64
	default:
		return kind
	}
}

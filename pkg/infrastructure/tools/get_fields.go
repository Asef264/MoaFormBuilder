package tools

import "reflect"

// GetFields returns an array of values out of a struct with fields all in the same type
func GetFields[T any](p T) []any {
	v := reflect.ValueOf(p)
	fields := v.NumField()
	var result []any
	for i := 0; i < fields; i++ {
		field := v.Field(i).Interface()
		result = append(result, field)
	}
	return result
}

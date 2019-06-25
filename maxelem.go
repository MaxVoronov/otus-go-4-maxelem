package maxelem

import (
	"errors"
	"reflect"
)

// Find max value in slice using custom comparator function
func FindMax(slice interface{}, comparator func(i, j int) bool) (interface{}, error) {
	valueType := reflect.ValueOf(slice)
	if valueType.Kind() != reflect.Slice {
		return nil, errors.New("passed not slice variable")
	}

	if valueType.Len() <= 0 {
		return nil, errors.New("passed empty slice")
	}

	maxValueIdx := 0
	for idx := 1; idx < valueType.Len(); idx++ {
		if comparator(maxValueIdx, idx) {
			maxValueIdx = idx
		}
	}

	return valueType.Index(maxValueIdx).Interface(), nil
}

// Find max integer in slice
func FindMaxInt(slice []int) (int, error) {
	result, err := FindMax(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	return result.(int), err
}

// Find max float in slice
func FindMaxFloat(slice []float32) (float32, error) {
	result, err := FindMax(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	return result.(float32), err
}

// Find max string in slice
func FindMaxString(slice []string) (string, error) {
	result, err := FindMax(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	return result.(string), err
}

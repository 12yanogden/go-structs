package structs

import "reflect"

func Equals(s1 interface{}, s2 interface{}) bool {
	r1 := reflect.ValueOf(s1)
	r2 := reflect.ValueOf(s2)

	if r1.NumField() != r2.NumField() {
		return false
	}

	for i := 0; i < r1.NumField(); i++ {
		if r1.Field(i).Interface() != r2.Field(i).Interface() {
			return false
		}
	}

	return true
}

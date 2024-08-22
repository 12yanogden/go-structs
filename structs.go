package structs

import "reflect"

func Equals(s1 interface{}, s2 interface{}) bool {
	r1 := reflect.ValueOf(s1)
	r2 := reflect.ValueOf(s2)
	// TODO: Need to return a list of issues as well
	// issues := []string{}

	if r1.NumField() != r2.NumField() {
		return false
	}

	for i := 0; i < r1.NumField(); i++ {
		f1 := r1.Field(i).Interface()
		f2 := r2.Field(i).Interface()
		// TODO: Need to compare field names
		// toString?
		// Might be able to just use == or DeeplyEqual()

		if f1 != f2 {
			return false
		}
	}

	return true
}

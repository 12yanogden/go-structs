package structs

import "reflect"

func Equals(s1 interface{}, s2 interface{}) bool {
	rv := reflect.ValueOf(s1)

	for i := 0; i < rv.NumField(); i++ {
		rvf := rv.Field(i)
		rfv := reflect.ValueOf(s2).Field(i)

		if rvf.Interface() != rfv.Interface() {
			return false
		}
	}

	return true
}

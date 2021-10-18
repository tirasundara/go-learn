package reflections

import "reflect"

func walk(x interface{}, fn func(input string)) {
	v := reflect.ValueOf(x)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fn(field.String())
	}
}

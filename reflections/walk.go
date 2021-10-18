package reflections

import "reflect"

func walk(x interface{}, fn func(input string)) {
	v := reflect.ValueOf(x)
	field := v.Field(0)
	fn(field.String())
}

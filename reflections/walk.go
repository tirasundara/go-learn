package reflections

import "reflect"

func walk(x interface{}, fn func(input string)) {
	v := getValue(x)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

func getValue(x interface{}) (val reflect.Value) {
	val = reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		return val.Elem()
	}

	return
}

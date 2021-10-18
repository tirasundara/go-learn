package reflections

import "reflect"

func walk(x interface{}, fn func(input string)) {
	v := getValue(x)

	if v.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			walk(v.Index(i).Interface(), fn)
		}

		return
	}

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

package reflections

import "reflect"

func walk(x interface{}, fn func(input string)) {
	v := getValue(x)

	switch v.Kind() {
	case reflect.String:
		fn(v.String())
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			walk(v.Field(i).Interface(), fn)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			walk(v.Index(i).Interface(), fn)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			walk(v.MapIndex(key).Interface(), fn)
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

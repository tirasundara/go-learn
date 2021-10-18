package reflections

import "reflect"

func walk(x interface{}, fn func(input string)) {
	v := getValue(x)

	walkValue := func(r reflect.Value) {
		walk(r.Interface(), fn)
	}

	switch v.Kind() {
	case reflect.String:
		fn(v.String())
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			walkValue(v.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			walkValue(v.Index(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			walkValue(v.MapIndex(key))
		}
	case reflect.Chan:
		for val, ok := v.Recv(); ok; val, ok = v.Recv() {
			walkValue(val)
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

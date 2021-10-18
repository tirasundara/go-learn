package reflections

import "reflect"

func walk(x interface{}, fn func(input string)) {
	v := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch v.Kind() {
	case reflect.String:
		fn(v.String())
	case reflect.Slice:
		numberOfValues = v.Len()
		getField = v.Index
	case reflect.Struct:
		numberOfValues = v.NumField()
		getField = v.Field
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) (val reflect.Value) {
	val = reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		return val.Elem()
	}

	return
}

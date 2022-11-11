package walk

import (
	"reflect"
)

func Walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	switch val.Kind() {

	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			Walk(val.Field(i).Interface(), fn)
		}

	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			Walk(val.Index(i).Interface(), fn)
		}

	case reflect.String:
		fn(val.String())
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}

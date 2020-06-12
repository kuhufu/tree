package default_value

import (
	"fmt"
	"reflect"
)

/*
* 设置默认值
 */

func DefaultValue(obj interface{}) {
	valueOf := reflect.ValueOf(obj).Elem()
	typeOf := valueOf.Type()

	if valueOf.Kind() != reflect.Struct {
		panic(fmt.Errorf("type of obj  must be struct, but actual is %v", valueOf.Kind()))
	}

	for i := 0; i < valueOf.NumField(); i++ {
		val := valueOf.Field(i)
		field := typeOf.Field(i)

		if defaultVal, ok := field.Tag.Lookup("default"); ok {
			if val.IsZero() {
				if err := setWithProperType(defaultVal, val, field); err != nil {
					panic(err)
				}
			}
		}
	}
}

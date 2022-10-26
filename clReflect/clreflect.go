package clReflect

import (
	"github.com/xiaolan580230/clUtil/clLog"
	"reflect"
)

func StructToMap(_ref interface{}) map[string]interface{} {

	_type := reflect.TypeOf(_ref)
	_value := reflect.ValueOf(_ref)

	var ma = make(map[string]interface{})

	for i := 0; i < _type.NumField(); i++ {
		if _type.Field(i).Name[0] >= 97 {
			continue	// 小写的过滤掉
		}
		clLog.Info("%+v Anonymous: %v", _type.Field(i), _type.Field(i).Anonymous)

		var fileName = _type.Field(i).Tag.Get("db")
		var fieldValue = _value.Field(i).Interface()

		if fileName == "" {
			continue
		}

		ma[fileName] = fieldValue
	}
	return ma
}
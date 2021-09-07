package clJson

import (
	"fmt"
	"testing"
)

func TestCreateBy(t *testing.T) {

	jsonObj := New([]byte( `{"a":1000, "b":2}` ))
	if jsonObj == nil {
		fmt.Printf("jsonObj解析错误!\n")
		return
	}
	fmt.Printf(">> a的值: %v\n", jsonObj.GetInt32("a"))
}


func TestJsonStream_GetArray(t *testing.T) {

	jsonObj := New([]byte( `{"a":[1,2,3,4,5,6]}`))
	if jsonObj == nil {
		fmt.Printf("jsonObj解析错误!\n")
		return
	}
	
	jsonArr := jsonObj.GetArray("a")
	jsonArr.Each(func(key int, value *JsonStream) bool {
		fmt.Printf("val: %v\n", value.ToStr())
		return true
	})

	fmt.Printf("数组: %+v\n", jsonArr.ToCustom())
}
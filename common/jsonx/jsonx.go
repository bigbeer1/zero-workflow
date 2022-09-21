package jsonx

import (
	"container/list"
	"encoding/json"
	"fmt"
)

func ToJSONStr(data interface{}) (string, error) {
	result, err := json.Marshal(data)
	return fmt.Sprintf("%s", result), err
}

// Str2Struct Str2Struct
// 字符串转对象
func Str2Struct(source string, destination interface{}) error {

	err := json.Unmarshal([]byte(source), destination)
	return err
}

// list对象转数组
func List2Array(list *list.List) []interface{} {
	var len = list.Len()
	if len == 0 {
		return nil
	}
	var arr []interface{}
	for e := list.Front(); e != nil; e = e.Next() {
		arr = append(arr, e.Value)
	}
	return arr
}

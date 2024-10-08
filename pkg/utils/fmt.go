package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: StructToMap
//@description: 利用反射将结构体转化为map
//@param: obj interface{}
//@return: map[string]interface{}

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Tag.Get("mapstructure") != "" {
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ArrayToString
//@description: 将数组格式化为字符串
//@param: array []interface{}
//@return: string

func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

func ArrayToInterfaceInt(array []int) []interface{} {
	splice := make([]interface{}, len(array))
	for k, v := range array {
		splice[k] = v
	}
	return splice
}

func StrToFloat64(str string) float64 {
	f, _ := strconv.ParseFloat(str, 10)
	return f
}
func StrToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

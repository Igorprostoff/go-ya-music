package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ConvertTrackIdToNumber(trackId string) (int, error) {

	return strconv.Atoi(strings.Split(trackId, ":")[0])

}

func InArray(val interface{}, array interface{}) bool {

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				return true
			}
		}
	}

	return false
}
func ToMap(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil
}

func MapKeysInStructKeys(m map[string]interface{}, s interface{}) bool {
	m2, err := ToMap(s, "json")
	if err != nil {
		fmt.Println("Unable to convert to struct")
		return false
	}
	for k := range m {
		i := false
		for k2 := range m2 {
			k2 = strings.ReplaceAll(k2, ",omitempty", "")
			if k2 == k {
				i = true
			}
		}
		if !i {
			return false
		}
	}
	return true
}

/*
var (
	INSERT = "INSERT"
	DELETE = "DELETE"
)

func add_delete(from_, to int) {

	operation =
		map[string]string{
			"op": DELETE, "from": strconv.Itoa(from_), "to": strconv.Itoa(to),
		}

	self.operations.append(operation)
	return self
}

var (

)
*/

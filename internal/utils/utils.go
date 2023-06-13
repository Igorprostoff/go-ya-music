package utils

import (
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

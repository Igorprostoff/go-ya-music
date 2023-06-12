package utils

import (
	"reflect"
	"strconv"
	"strings"
)

func Convert_track_id_to_number(track_id string) (int, error) {

	return strconv.Atoi(strings.Split(track_id, ":")[0])

}

func In_array(val interface{}, array interface{}) bool {

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

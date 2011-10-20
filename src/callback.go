package curl

/*
#include <curl/curl.h>
*/
import "C"

import (
 	"unsafe"
 	"reflect"
)



//export getCurlField
func getCurlField(p uintptr, name string) uintptr {
	curl := (* CURL)(unsafe.Pointer(p))
	switch name {
	case "onDataAvailable":
		return reflect.ValueOf(curl.onDataAvailable).Pointer()
	case "onHeaderAvailable":
		return reflect.ValueOf(curl.onHeaderAvailable).Pointer()
	case "onProgressAvailable":
		return reflect.ValueOf(curl.onProgressAvailable).Pointer()
	}
	return 0
}

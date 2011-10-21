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
func getCurlField(p uintptr, cname *C.char) uintptr {
	name := C.GoString(cname)
	curl := (* CURL)(unsafe.Pointer(p))
	switch name {
	case "readFunction":
		return reflect.ValueOf(curl.readFunction).Pointer()
	case "headerFunction":
		return reflect.ValueOf(curl.headerFunction).Pointer()
	case "writeFunction":
		return reflect.ValueOf(curl.writeFunction).Pointer()
	case "progressFuncion":
		return reflect.ValueOf(curl.progressFuncion).Pointer()
	case "headerData":
		return uintptr(unsafe.Pointer(curl.headerData))
	case "writeData":
		return uintptr(unsafe.Pointer(curl.writeData))
	case "readData":
		return uintptr(unsafe.Pointer(curl.readData))
	}

	println("WARNING: field not found: ", name)
	return 0
}

//export nilInterface
func nilInterface() interface{}{
	return nil
}

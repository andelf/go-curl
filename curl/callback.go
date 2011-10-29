package curl

/*
#cgo linux pkg-config: libcurl
#include <stdlib.h>
#include <string.h>
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
	case "progressFunction":
		return reflect.ValueOf(curl.progressFunction).Pointer()
	case "headerData":
		return uintptr(unsafe.Pointer(curl.headerData))
	case "writeData":
		return uintptr(unsafe.Pointer(curl.writeData))
	case "readData":
		return uintptr(unsafe.Pointer(curl.readData))
	case "progressData":
		return uintptr(unsafe.Pointer(curl.progressData))
	}

	println("WARNING: field not found: ", name)
	return 0
}

//export nilInterface
func nilInterface() interface{}{
	return nil
}

// callback functions
//export callWriteFunctionCallback
func callWriteFunctionCallback(
	f func([]byte, uintptr, interface{}) uintptr,
	ptr *C.char,
	size C.size_t,
	userdata interface{}) uintptr {
	buf := C.GoBytes(unsafe.Pointer(ptr), C.int(size))
	ret := f(buf, uintptr(size), userdata)
	return ret
}

//export callProgressCallback
func callProgressCallback(
	f func(interface{}, float64, float64, float64, float64) int,
	clientp interface{},
	dltotal, dlnow, ultotal, ulnow C.double) int {
	// fdltotal, fdlnow, fultotal, fulnow
	ret := f(clientp, float64(dltotal), float64(dlnow), float64(ultotal), float64(ulnow))
	return ret
}

//export callReadFunctionCallback
func callReadFunctionCallback(
	f func([]byte, uintptr, interface{}) uintptr,
	ptr *C.char,
	size C.size_t,
	userdata interface{}) uintptr {
	// TODO code cleanup
	buf := C.GoBytes(unsafe.Pointer(ptr), C.int(size))
	ret := f(buf, uintptr(size), userdata)
	str := C.CString(string(buf))
	defer C.free(unsafe.Pointer(str))
	if C.memcpy(unsafe.Pointer(ptr), unsafe.Pointer(str), C.size_t(ret)) == nil {
		panic("read_callback memcpy error!")
	}
	return ret
}

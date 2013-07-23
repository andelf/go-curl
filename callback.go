package curl

/*
#include <stdlib.h>
#include <string.h>
#include <curl/curl.h>

*/
import "C"

import (
	"reflect"
	"unsafe"
)

//export goGetCurlField
func goGetCurlField(p uintptr, cname *C.char) uintptr {
	name := C.GoString(cname)
	curl := (*CURL)(unsafe.Pointer(p))
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

	warnf("Field not found: %s", name)
	return 0
}

//export goNilInterface
func goNilInterface() interface{} {
	return nil
}

// callback functions
//export goCallWriteFunctionCallback
func goCallWriteFunctionCallback(f *func([]byte, interface{}) bool,
	ptr *C.char,
	size C.size_t,
	userdata interface{}) uintptr {
	buf := C.GoBytes(unsafe.Pointer(ptr), C.int(size))
	ok := (*f)(buf, userdata)
	if ok {
		return uintptr(size)
	}
	//return uintptr(C.CURL_MAX_WRITE_SIZE + 1)
	return C.CURL_WRITEFUNC_PAUSE
}

//export goCallProgressCallback
func goCallProgressCallback(f *func(float64, float64, float64, float64, interface{}) bool,
	userdata interface{},
	dltotal, dlnow, ultotal, ulnow C.double) int {
	// fdltotal, fdlnow, fultotal, fulnow
	ok := (*f)(float64(dltotal), float64(dlnow), float64(ultotal), float64(ulnow), userdata)
	// non-zero va lue will cause libcurl to abort the transfer and return Error
	if ok {
		return 0
	}
	return 1
}

//export goCallReadFunctionCallback
func goCallReadFunctionCallback(f *func([]byte, interface{}) int,
	ptr *C.char,
	size C.size_t,
	userdata interface{}) uintptr {
	// TODO code cleanup
	buf := C.GoBytes(unsafe.Pointer(ptr), C.int(size))
	ret := (*f)(buf, userdata)
	str := C.CString(string(buf))
	defer C.free(unsafe.Pointer(str))
	if C.memcpy(unsafe.Pointer(ptr), unsafe.Pointer(str), C.size_t(ret)) == nil {
		panic("read_callback memcpy error!")
	}
	return uintptr(ret)
}

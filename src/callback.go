package curl

/*
#include <curl/curl.h>
*/
import "C"

// import (
// 	"unsafe"
// 	"reflect"
// )



type ReadFunc func(ptr []byte, size, nmemb uintptr, userdata interface{}) uintptr

//export callReadFunc
func callReadFunc(f interface{}, ptr interface{}, size, nmemb uintptr, userdata interface{}) uintptr {
	fun := f.(ReadFunc)
	bufptr := ptr.([]byte)

	return fun(bufptr, size, nmemb, userdata)
}

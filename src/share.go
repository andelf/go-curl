
package curl

/*
#include <curl/curl.h>
*/
import "C"
import (
	"unsafe"
)

/*
TODO
 CURLSHcode curl_share_setopt(CURLSH *, CURLSHoption option, ...);
*/

type SHCode C.CURLSHcode

func (errornum SHCode) String() string {
	// ret is const char*, no need to free
	ret := C.curl_share_strerror(C.CURLSHcode(errornum))
	return C.GoString(ret)
}

type CURLSH struct {
	handle unsafe.Pointer
}

func ShareInit() *CURLSH {
	p := C.curl_share_init()
	return &CURLSH{p}
}

func (shcurl *CURLSH) Cleanup() MCode {
	p := shcurl.handle
	return MCode(C.curl_share_cleanup(p))
}

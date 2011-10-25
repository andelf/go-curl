
package curl


// #cgo linux pkg-config: libcurl
// #include <curl/curl.h>
import "C"
import (
	"unsafe"
	"os"
)

// implement os.Error interface
type CurlShareError C.CURLMcode

func (e CurlShareError) String() string {
	// ret is const char*, no need to free
	ret := C.curl_share_strerror(C.CURLSHcode(e))
	return C.GoString(ret)
}


func newCurlShareError(errno C.CURLSHcode) os.Error {
	if errno == C.CURLSHE_OK {		// if nothing wrong
		return nil
	}
	return CurlShareError(errno)
}


type CURLSH struct {
	handle unsafe.Pointer
}

func ShareInit() *CURLSH {
	p := C.curl_share_init()
	return &CURLSH{p}
}

func (shcurl *CURLSH) Cleanup() os.Error {
	p := shcurl.handle
	return newCurlShareError(C.curl_share_cleanup(p))
}

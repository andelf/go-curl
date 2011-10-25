package curl



// #cgo linux pkg-config: libcurl
// #include <stdlib.h>
// #include <curl/curl.h>
import "C"
import (
	"unsafe"
	"os"
)

type CurlMultiError C.CURLMcode

func (e CurlMultiError) String() string {
	// ret is const char*, no need to free
	ret := C.curl_multi_strerror(C.CURLMcode(e))
	return C.GoString(ret)
}


func newCurlMultiError(errno C.CURLMcode) os.Error {
	if errno == C.CURLM_OK {		// if nothing wrong
		return nil
	}
	return CurlMultiError(errno)
}


type CURLM struct {
	handle unsafe.Pointer
}

// ok
func MultiInit() *CURLM {
	p := C.curl_multi_init()
	return &CURLM{p}
}

// ok
func (mcurl *CURLM) Cleanup() os.Error {
	p := mcurl.handle
	return newCurlMultiError(C.curl_multi_cleanup(p))
}

// ok
func (mcurl *CURLM) Perform() (int, os.Error) {
	p := mcurl.handle
	running_handles := C.int(-1)
	err := newCurlMultiError(C.curl_multi_perform(p, &running_handles))
	return int(running_handles), err
}

// ok
func (mcurl *CURLM) AddHandle(easy *CURL) os.Error {
	mp := mcurl.handle
	easy_handle := easy.handle
	return newCurlMultiError(C.curl_multi_add_handle(mp, easy_handle))
}

func (mcurl *CURLM) RemoveHandle(easy *CURL) os.Error {
	mp := mcurl.handle
	easy_handle := easy.handle
	return newCurlMultiError(C.curl_multi_remove_handle(mp, easy_handle))
}

func (mcurl *CURLM) Timeout() (int, os.Error) {
	p := mcurl.handle
	timeout := C.long(-1)
	err := newCurlMultiError(C.curl_multi_timeout(p, &timeout))
	return int(timeout), err
}


// TODO
/*
 curl_multi_info_read      curl_multi_setopt
 curl_multi_assign                    !curl_multi_socket
 curl_multi_socket_action
 curl_multi_fdset            ?curl_multi_strerror
*/

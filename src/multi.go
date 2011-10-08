package curl


/*
#include <stdlib.h>
#include <curl/curl.h>
*/
import "C"
import (
	"unsafe"
)

type MCode C.CURLMcode

func (errornum MCode) String() string {
	// ret is const char*, no need to free
	ret := C.curl_multi_strerror(C.CURLMcode(errornum))
	return C.GoString(ret)
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
func (mcurl *CURLM) Cleanup() MCode {
	p := mcurl.handle
	return MCode(C.curl_multi_cleanup(p))
}

// ok
func (mcurl *CURLM) Perform() (MCode, int) {
	p := mcurl.handle
	running_handles := C.int(-1)
	ret := MCode(C.curl_multi_perform(p, &running_handles))
	return ret, int(running_handles)
}

// ok
func (mcurl *CURLM) AddHandle(easy *CURL) MCode {
	mp := mcurl.handle
	easy_handle := easy.handle
	return MCode(C.curl_multi_add_handle(mp, easy_handle))
}

func (mcurl *CURLM) RemoveHandle(easy *CURL) MCode {
	mp := mcurl.handle
	easy_handle := easy.handle
	return MCode(C.curl_multi_remove_handle(mp, easy_handle))
}

func (mcurl *CURLM) Timeout() (MCode, int) {
	p := mcurl.handle
	timeout := C.long(-1)
	ret := MCode(C.curl_multi_timeout(p, &timeout))
	return ret, int(timeout)
}


// TODO
/*
 curl_multi_info_read      curl_multi_setopt
 curl_multi_assign                    !curl_multi_socket
 curl_multi_socket_action
 curl_multi_fdset            ?curl_multi_strerror
*/

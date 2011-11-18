package curl

/*
#cgo linux pkg-config: libcurl
#include <stdlib.h>
#include <curl/curl.h>
static CURLMcode curl_multi_setopt_long(CURLM *handle, CURLMoption option, long parameter) {
  return curl_multi_setopt(handle, option, parameter);
}
static CURLMcode curl_multi_setopt_pointer(CURLM *handle, CURLMoption option, void *parameter) {
  return curl_multi_setopt(handle, option, parameter);
}
*/
import "C"

import "unsafe"


type CurlMultiError C.CURLMcode

func (e CurlMultiError) Error() string {
	// ret is const char*, no need to free
	ret := C.curl_multi_strerror(C.CURLMcode(e))
	return C.GoString(ret)
}

func newCurlMultiError(errno C.CURLMcode) error {
	if errno == C.CURLM_OK { // if nothing wrong
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
func (mcurl *CURLM) Cleanup() error {
	p := mcurl.handle
	return newCurlMultiError(C.curl_multi_cleanup(p))
}

// ok
func (mcurl *CURLM) Perform() (int, error) {
	p := mcurl.handle
	running_handles := C.int(-1)
	err := newCurlMultiError(C.curl_multi_perform(p, &running_handles))
	return int(running_handles), err
}

// ok
func (mcurl *CURLM) AddHandle(easy *CURL) error {
	mp := mcurl.handle
	easy_handle := easy.handle
	return newCurlMultiError(C.curl_multi_add_handle(mp, easy_handle))
}

func (mcurl *CURLM) RemoveHandle(easy *CURL) error {
	mp := mcurl.handle
	easy_handle := easy.handle
	return newCurlMultiError(C.curl_multi_remove_handle(mp, easy_handle))
}

func (mcurl *CURLM) Timeout() (int, error) {
	p := mcurl.handle
	timeout := C.long(-1)
	err := newCurlMultiError(C.curl_multi_timeout(p, &timeout))
	return int(timeout), err
}

func (mcurl *CURLM) Setopt(opt int, param interface{}) error {
	p := mcurl.handle
	if param == nil {
		return newCurlMultiError(C.curl_multi_setopt_pointer(p, C.CURLMoption(opt), nil))
	}
	switch opt {
	//  currently cannot support these option
	//	case MOPT_SOCKETFUNCTION, MOPT_SOCKETDATA, MOPT_TIMERFUNCTION, MOPT_TIMERDATA:
	//		panic("not supported CURLM.Setopt opt")
	case MOPT_PIPELINING, MOPT_MAXCONNECTS:
		val := C.long(0)
		switch t := param.(type) {
		case int:
			val := C.long(t)
			return newCurlMultiError(C.curl_multi_setopt_long(p, C.CURLMoption(opt), val))
		case bool:
			val = C.long(0)
			if t {
				val = C.long(1)
			}
			return newCurlMultiError(C.curl_multi_setopt_long(p, C.CURLMoption(opt), val))
		}
	}
	panic("not supported CURLM.Setopt opt or param")
	return nil
}

// TODO curl_multi_info_read

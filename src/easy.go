package curl

/*
 #cgo linux pkg-config: libcurl
#include <stdlib.h>
#include <curl/curl.h>
#include "callback.h"
static CURLcode curl_easy_setopt_long(CURL *handle, CURLoption option, long parameter) {
  return curl_easy_setopt(handle, option, parameter);
}
static CURLcode curl_easy_setopt_string(CURL *handle, CURLoption option, char *parameter) {
  return curl_easy_setopt(handle, option, parameter);
}
static CURLcode curl_easy_setopt_slist(CURL *handle, CURLoption option, struct curl_slist *parameter) {
  return curl_easy_setopt(handle, option, parameter);
}
static CURLcode curl_easy_setopt_pointer(CURL *handle, CURLoption option, void *parameter) {
  return curl_easy_setopt(handle, option, parameter);
}

static CURLcode curl_easy_getinfo_string(CURL *curl, CURLINFO info, char **p) {
 return curl_easy_getinfo(curl, info, p);
}
static CURLcode curl_easy_getinfo_long(CURL *curl, CURLINFO info, long *p) {
 return curl_easy_getinfo(curl, info, p);
}
static CURLcode curl_easy_getinfo_double(CURL *curl, CURLINFO info, double *p) {
 return curl_easy_getinfo(curl, info, p);
}
static CURLcode curl_easy_getinfo_slist(CURL *curl, CURLINFO info, struct curl_slist *p) {
 return curl_easy_getinfo(curl, info, p);
}
*/
import "C"

import (
	"unsafe"
	"reflect"
	"os"
	"fmt"
)

/*
*/
// consts
/*
const (
	OPTTYPE_LONG          = 0
	OPTTYPE_OBJECTPOINT   = 10000
	OPTTYPE_FUNCTIONPOINT = 20000
	OPTTYPE_OFF_T         = 30000
)
*/

const (
	PAUSE_RECV      = C.CURLPAUSE_RECV
	PAUSE_RECV_CONT = C.CURLPAUSE_RECV_CONT
	PAUSE_SEND      = C.CURLPAUSE_SEND
	PAUSE_SEND_CONT = C.CURLPAUSE_SEND_CONT
	PAUSE_ALL       = C.CURLPAUSE_ALL
	PAUSE_CONT      = C.CURLPAUSE_CONT
)

// ======================== functions ========



// all ret code
type CurlError C.CURLcode

func (e CurlError) String() string {
	// ret is const char*, no need to free
	ret := C.curl_easy_strerror(C.CURLcode(e))
	return fmt.Sprintf("curl: %s", C.GoString(ret))
}

func newCurlError(errno C.CURLcode) os.Error {
	if errno == C.CURLE_OK {		// if nothing wrong
		return nil
	}
	return CurlError(errno)
}

// curl_easy interface
type CURL struct {
	handle unsafe.Pointer
	headerFunction, writeFunction, readFunction func([]byte, uintptr, interface{}) uintptr
	progressFunction func(interface{}, float64, float64, float64, float64) int
	fnmatchFunction func(interface{}, string, string) int
	headerData, writeData, readData, progressData, fnmatchData *interface{}
}

func EasyInit() *CURL {
	p := C.curl_easy_init()
	return &CURL{handle: p}		// other field defaults to nil
}

func (curl *CURL) Duphandle() *CURL {
	p := curl.handle
	return &CURL{handle: C.curl_easy_duphandle(p)}
}

func (curl *CURL) Cleanup() {
	p := curl.handle
	C.curl_easy_cleanup(p)
}


// WARNING: why ? function pointer is &fun, but function addr is reflect.ValueOf(fun).Pointer()
func (curl *CURL) Setopt(opt int, param interface{}) os.Error {
	p := curl.handle
	if param == nil {
		// NOTE: some option will crash program when got a nil param
		return newCurlError(C.curl_easy_setopt_pointer(p, C.CURLoption(opt), nil))
	}
	switch {
	case opt == OPT_READDATA:
		curl.readData = &param
		return nil
	case opt == OPT_READFUNCTION:
		fun := param.(func([]byte, uintptr, interface{}) uintptr)
		curl.readFunction = fun

		ptr := C.return_read_function()
		if err := newCurlError(C.curl_easy_setopt_pointer(p, C.CURLoption(opt), ptr)); err == nil {
			return newCurlError(C.curl_easy_setopt_pointer(p, OPT_READDATA,
				unsafe.Pointer(reflect.ValueOf(curl).Pointer())))
		} else {
			return err
		}

	case opt == OPT_PROGRESSDATA:
		curl.progressData = &param
		return nil
	case opt == OPT_PROGRESSFUNCTION:
		fun := param.(func(interface{}, float64, float64, float64, float64) int)
		curl.progressFunction = fun

		ptr := C.return_progress_function()
		if err := newCurlError(C.curl_easy_setopt_pointer(p, C.CURLoption(opt), ptr)); err == nil {
			return newCurlError(C.curl_easy_setopt_pointer(p, OPT_PROGRESSDATA,
				unsafe.Pointer(reflect.ValueOf(curl).Pointer())))
		} else {
			return err
		}

	case opt == OPT_HEADERDATA:	// also known as OPT_WRITEHEADER
		curl.headerData = &param
		return nil
	case opt == OPT_HEADERFUNCTION:
		fun := param.(func([]byte, uintptr, interface{}) uintptr)
		curl.headerFunction = fun

		ptr := C.return_header_function()
		if err := newCurlError(C.curl_easy_setopt_pointer(p, C.CURLoption(opt), ptr)); err == nil {
			return newCurlError(C.curl_easy_setopt_pointer(p, OPT_HEADERDATA,
				unsafe.Pointer(reflect.ValueOf(curl).Pointer())))
		} else {
			return err
		}
	// just copy & modification of above
	case opt == OPT_WRITEDATA:
	 	curl.writeData = &param
		return nil
	case opt == OPT_WRITEFUNCTION:
		fun := param.(func([]byte, uintptr, interface{}) uintptr)
		curl.writeFunction = fun

		ptr := C.return_write_function()
		if err := newCurlError(C.curl_easy_setopt_pointer(p, C.CURLoption(opt), ptr)); err == nil {
			return newCurlError(C.curl_easy_setopt_pointer(p, OPT_WRITEDATA,
				unsafe.Pointer(reflect.ValueOf(curl).Pointer())))
		} else {
			return err
		}

	case opt > C.CURLOPTTYPE_OFF_T:
		// here we should use uint64
		panic("off_t type not implemented yet!")
	case opt > C.CURLOPTTYPE_FUNCTIONPOINT:
		// function pointer
		panic("function poionter not implemented yet!")

	case opt > C.CURLOPTTYPE_OBJECTPOINT:
		switch t := param.(type) {
		case string:
			// FIXME: memory leak, some opt needs we hold a c string till perform()
			// TODO: We can add a []unsafe.Poionter to Curl struct and do cleanup in Cleanup()
			ptr := C.CString(t)
			// defer C.free(unsafe.Pointer(ptr))
			ret := C.curl_easy_setopt_string(p, C.CURLoption(opt), ptr)
			return newCurlError(ret)
		case []string:
			if len(t) > 0 {
				a_slist := C.curl_slist_append(nil, C.CString(t[0]))
				for _, s := range t[1:] {
					a_slist = C.curl_slist_append(a_slist, C.CString(s))
				}
				return newCurlError(C.curl_easy_setopt_slist(p, C.CURLoption(opt), a_slist))
			} else {
				return newCurlError(C.curl_easy_setopt_slist(p, C.CURLoption(opt), nil))
			}
		default:
			// It panics if v's Kind is not Chan, Func, Map, Ptr, Slice, or UnsafePointer.
			// val := reflect.ValueOf(param)
			//fmt.Printf("DEBUG(Setopt): param=%x\n", val.Pointer())
			//println("DEBUG can addr =", val.Pointer(), "opt=", opt)
			// pass a pointer to GoInterface
			return newCurlError(C.curl_easy_setopt_pointer(p, C.CURLoption(opt),
				unsafe.Pointer(&param)))
		}
	case opt > C.CURLOPTTYPE_LONG:
		// long
		switch t := param.(type) {
		case int:
			val := C.long(t)
			ret := C.curl_easy_setopt_long(p, C.CURLoption(opt), val)
			return newCurlError(ret)
		case bool:
			val := 0
			if t {
				val = 1
			}
			ret := C.curl_easy_setopt_long(p, C.CURLoption(opt), C.long(val))
			return newCurlError(ret)
		default:
			panic("type error in param")
		}
	}
	panic("opt param error!")
}

func (curl *CURL) Send(buffer []byte) (int, os.Error) {
	p := curl.handle
	buflen := len(buffer)
	n := C.size_t(0)
	ret := C.curl_easy_send(p, unsafe.Pointer(&buffer[0]), C.size_t(buflen), &n)
	return int(n), newCurlError(ret)
}

func (curl *CURL) Recv(buffer []byte) (int, os.Error) {
	p := curl.handle
	buflen := len(buffer)
	buf := C.CString(string(buffer))
	n := C.size_t(0)
	ret := C.curl_easy_recv(p, unsafe.Pointer(buf), C.size_t(buflen), &n)
	return copy(buffer, C.GoStringN(buf, C.int(n))), newCurlError(ret)

}

func (curl *CURL) Perform() os.Error {
	p := curl.handle
	return newCurlError(C.curl_easy_perform(p))
}

func (curl *CURL) Pause(bitmask int) os.Error {
	p := curl.handle
	return newCurlError(C.curl_easy_pause(p, C.int(bitmask)))
}

func (curl *CURL) Reset() {
	p := curl.handle
	C.curl_easy_reset(p)
}

func (curl *CURL) Escape(url string) string {
	p := curl.handle
	oldUrl := C.CString(url)
	defer C.free(unsafe.Pointer(oldUrl))
	newUrl := C.curl_easy_escape(p, oldUrl, 0)
	defer C.curl_free(unsafe.Pointer(newUrl))
	return C.GoString(newUrl)
}

func (curl *CURL) Unescape(url string) string {
	p := curl.handle
	oldUrl := C.CString(url)
	outlength := C.int(0)
	defer C.free(unsafe.Pointer(oldUrl))
	// If outlength is non-NULL, the function will write the length of the
	// returned string in  the  integer  it  points  to.  This allows an
	// escaped string containing %00 to still get used properly after unescaping.
	newUrl := C.curl_easy_unescape(p, oldUrl, 0, &outlength)
	defer C.curl_free(unsafe.Pointer(newUrl))
	return C.GoStringN(newUrl, outlength)
}

/*
 CURLINFO_STRING   0x100000
 CURLINFO_LONG     0x200000
 CURLINFO_DOUBLE   0x300000
 CURLINFO_SLIST    0x400000
 CURLINFO_MASK     0x0fffff
 CURLINFO_TYPEMASK 0xf00000
 */

const (
	_INFO_STRING = C.CURLINFO_STRING
	_INFO_LONG = C.CURLINFO_LONG
	_INFO_DOUBLE = C.CURLINFO_DOUBLE
	_INFO_SLIST = C.CURLINFO_SLIST
	_INFO_MASK = C.CURLINFO_MASK
	_INFO_TYPEMASK = C.CURLINFO_TYPEMASK
)

func (curl *CURL) Getinfo(info C.CURLINFO) (ret interface{}, err os.Error) {
	p := curl.handle
	switch info & _INFO_TYPEMASK {
	case _INFO_STRING:
		a_string := C.CString("")
		defer C.free(unsafe.Pointer(a_string))
		err := newCurlError(C.curl_easy_getinfo_string(p, info, &a_string));
		ret := C.GoString(a_string)
		print("debug (Getinfo) ", ret, "\n")
		return ret, err
	case _INFO_LONG:
		a_long := C.long(-1)
		err := newCurlError(C.curl_easy_getinfo_long(p, info, &a_long));
		ret := int(a_long)
		print("debug (Getinfo) ", ret, "\n")
		return ret, err
	case _INFO_DOUBLE:
		a_double := C.double(0.0)
		err := newCurlError(C.curl_easy_getinfo_double(p, info, &a_double));
		ret := float64(a_double)
		print("debug (Getinfo) ", ret, "\n")
		return ret, err
	case _INFO_SLIST:			// need fix
		a_ptr_slist := new(_Ctype_struct_curl_slist)
		err := newCurlError(C.curl_easy_getinfo_slist(p, info, a_ptr_slist));
		ret := []string{}
		for a_ptr_slist != nil {
			print("!!debug (Getinfo) ", C.GoString(a_ptr_slist.data), a_ptr_slist.next, "\n")
			ret = append(ret, C.GoString(a_ptr_slist.data))
			a_ptr_slist = a_ptr_slist.next
		}
		return ret, err
	default:
		panic("error calling Getinfo\n")
	}
	panic("not implemented yet!")
	return nil, nil
}

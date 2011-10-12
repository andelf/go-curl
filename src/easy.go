package curl

/*
#include <stdlib.h>
#include <curl/curl.h>
static CURLcode curl_easy_setopt_long(CURL *handle, CURLoption option, long parameter) {
  return curl_easy_setopt(handle, option, parameter);
}
static CURLcode curl_easy_setopt_string(CURL *handle, CURLoption option, char *parameter) {
  return curl_easy_setopt(handle, option, parameter);
}
static CURLcode curl_easy_setopt_slist(CURL *handle, CURLoption option, struct curl_slist *parameter) {
  return curl_easy_setopt(handle, option, parameter);
}
// get info
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
)

/*
static CURLcode curl_easy_setopt_pointer(CURL *handle, CURLoption option, void *parameter) {
  return curl_easy_setopt(handle, option, parameter);
}
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
type Code C.CURLcode

func (errornum Code) String() string {
	// ret is const char*, no need to free
	ret := C.curl_easy_strerror(C.CURLcode(errornum))
	return C.GoString(ret)
}


type CURL struct {
	handle unsafe.Pointer
}

// easy handle
func EasyInit() *CURL {
	p := C.curl_easy_init()
	return &CURL{p}
}

func (curl *CURL) Duphandle() *CURL {
	p := curl.handle
	return &CURL{C.curl_easy_duphandle(p)}
}

func (curl *CURL) Cleanup() {
	p := curl.handle
	C.curl_easy_cleanup(p)
}

func (curl *CURL) Perform() Code {
	p := curl.handle
	return Code(C.curl_easy_perform(p))
}

func (curl *CURL) Setopt(opt int, param interface{}) Code {
	p := curl.handle
	// C.CURLoption
	switch {
	case opt > C.CURLOPTTYPE_OFF_T:
		//
		println("> off_t")
		break
	case opt > C.CURLOPTTYPE_FUNCTIONPOINT:
		// function pointer
		break
	case opt > C.CURLOPTTYPE_OBJECTPOINT:
		switch t := param.(type) {
		case string:
			// FIXME: memory leak, some opt needs
			ptr := C.CString(t)
			// defer C.free(unsafe.Pointer(ptr))
			ret := C.curl_easy_setopt_string(p, C.CURLoption(opt), ptr)
			return Code(ret)
		case []string:
			print("my debug =>", "creating a list")
			if len(t) > 0 {
				a_slist := C.curl_slist_append(nil, C.CString(t[0]))
				for _, s := range t[1:] {
					a_slist = C.curl_slist_append(a_slist, C.CString(s))
				}
				return Code(C.curl_easy_setopt_slist(p, C.CURLoption(opt), a_slist))
			} else {
				return Code(C.curl_easy_setopt_slist(p, C.CURLoption(opt), nil))
			}
		default:
			val := reflect.ValueOf(param)
			if val.CanAddr() {
				//println(val)
				println("=>", val.Addr().Pointer())
				ret := C.curl_easy_setopt_long(p, C.CURLoption(opt),
					C.long(val.Addr().Pointer()))
				println(ret)
			} else {
				panic("type error in param")
				return Code(1)
			}
		}
	case opt > C.CURLOPTTYPE_LONG:
		// long
		switch t := param.(type) {
		case int:
			val := C.long(t)
			ret := C.curl_easy_setopt_long(p, C.CURLoption(opt), val)
			return Code(ret)
		case bool:
			val := 0
			if t {
				val = 1
			}
			ret := C.curl_easy_setopt_long(p, C.CURLoption(opt), C.long(val))
			return Code(ret)
		default:
			panic("type error in param")
			return Code(1)
		}
	default:
		panic("opt param error!")
		return Code(1)
	}
	return Code(1)
}

// TODO: curl_easy_recv
// TODO: curl_easy_send


func (curl *CURL) Pause(bitmask int) Code {
	p := curl.handle
	return Code(C.curl_easy_pause(p, C.int(bitmask)))
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

func (curl *CURL) Getinfo(info C.CURLINFO) (Code, interface{}) {
	p := curl.handle
	switch info & _INFO_TYPEMASK {
	case _INFO_STRING:
		a_string := C.CString("")
		defer C.free(unsafe.Pointer(a_string))
		ret := Code(C.curl_easy_getinfo_string(p, info, &a_string));
		print("debug (Getinfo) ", C.GoString(a_string), "\n")
		return ret, C.GoString(a_string)
	case _INFO_LONG:
		a_long := C.long(-1)
		ret := Code(C.curl_easy_getinfo_long(p, info, &a_long));
		print("debug (Getinfo) ", int(a_long), "\n")
		return ret, int(a_long)
	case _INFO_DOUBLE:
		a_double := C.double(0.0)
		ret := Code(C.curl_easy_getinfo_double(p, info, &a_double));
		print("debug (Getinfo) ", float64(a_double), "\n")
		return ret, float64(a_double)
	case _INFO_SLIST:			// need fix
		a_ptr_slist := new(_Ctype_struct_curl_slist)
		ret := Code(C.curl_easy_getinfo_slist(p, info, a_ptr_slist));
		ret_slist := []string{}
		for a_ptr_slist != nil {
			print("!!debug (Getinfo) ", C.GoString(a_ptr_slist.data), a_ptr_slist.next, "\n")
			ret_slist = append(ret_slist, C.GoString(a_ptr_slist.data))
			a_ptr_slist = a_ptr_slist.next
		}
		return ret, ret_slist
	default:
		panic("error calling Getinfo\n")
	}
	println("Not implemented yet.")
	return Code(100), 0
}

// func EasyStrerror(errornum Code) string {
// 	ret := C.curl_easy_strerror(C.CURLcode(errornum))
// 	return C.GoString(ret)

// }

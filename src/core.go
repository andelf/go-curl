package curl

/*
#include <stdlib.h>
#include <curl/curl.h>
static char *string_array_index(char **p, int i) {
  return p[i];
}
*/
import "C"

import (
	"unsafe"
	"time"
	"os"
)

const (
	GLOBAL_SSL = C.CURL_GLOBAL_SSL
	GLOBAL_WIN32 = C.CURL_GLOBAL_WIN32
	GLOBAL_ALL = C.CURL_GLOBAL_ALL
	GLOBAL_NOTHING = C.CURL_GLOBAL_NOTHING
	GLOBAL_DEFAULT = C.CURL_GLOBAL_DEFAULT
)


func GlobalInit(flags int) os.Error {
	return newCurlError(C.curl_global_init(C.long(flags)))
}

// TODO: curl_global_init_mem
func GlobalInitMem(args ...interface{}) {
	panic("curl_global_init_mem not implemented yet!")
}


func GlobalCleanup() {
	C.curl_global_cleanup()
}

type VersionInfoData struct {
	Age C.CURLversion
	// age >= 0
	Version       string
	VersionNum    uint
	Host          string
	Features      int
	SslVersion    string
	SslVersionNum int
	LibzVersion   string
	Protocols     []string
	// age >= 1
	Ares    string
	AresNum int
	// age >= 2
	Libidn string
	// age >= 3
	IconvVerNum   int
	LibsshVersion string
}

const (
	VERSION_FIRST  = C.CURLVERSION_FIRST
	VERSION_SECOND = C.CURLVERSION_SECOND
	VERSION_THIRD  = C.CURLVERSION_THIRD
	VERSION_FOURTH = C.CURLVERSION_FOURTH
	VERSION_LAST   = C.CURLVERSION_LAST
	VERSION_NOW    = C.CURLVERSION_NOW
)


func Version() string {
	return C.GoString(C.curl_version())
}

// curl_version_info ok
func VersionInfo(ver C.CURLversion) *VersionInfoData {
	data := C.curl_version_info(ver)
	ret := new(VersionInfoData)
	ret.Age = data.age
	switch age := ret.Age; {
	case age >= 0:
		ret.Version = string(C.GoString(data.version))
		ret.VersionNum = uint(data.version_num)
		ret.Host = C.GoString(data.host)
		ret.Features = int(data.features)
		ret.SslVersion = C.GoString(data.ssl_version)
		ret.SslVersionNum = int(data.ssl_version_num)
		ret.LibzVersion = C.GoString(data.libz_version)
		// ugly but works
		ret.Protocols = []string{}
		for i := C.int(0); C.string_array_index(data.protocols, i) != nil; i++ {
			p := C.string_array_index(data.protocols, i)
			ret.Protocols = append(ret.Protocols, C.GoString(p))
		}
		fallthrough
	case age >= 1:
		ret.Ares = C.GoString(data.ares)
		ret.AresNum = int(data.ares_num)
		fallthrough
	case age >= 2:
		ret.Libidn = C.GoString(data.libidn)
		fallthrough
	case age >= 3:
		ret.IconvVerNum = int(data.iconv_ver_num)
		ret.LibsshVersion = C.GoString(data.libssh_version)
	}
	return ret
}


func Getdate(date string) *time.Time {
	datestr := C.CString(date)
	defer C.free(unsafe.Pointer(datestr))
	t := C.curl_getdate(datestr, nil)
	if t == -1 {
		return nil
	}
	return time.SecondsToUTC(int64(t))
}

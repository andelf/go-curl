package main

/*
#include <stdlib.h>
#include <sys/select.h>
static void _FD_ZERO(void *set) {
    FD_ZERO((fd_set*)set);
}
static void _FD_SET(int sysfd, void *set) {
    FD_SET(sysfd, (fd_set*)set);
}
static int _FD_ISSET (int sysfd, void *set) {
    return FD_ISSET(sysfd, (fd_set*)set);
}
*/
import "C"

import (
	curl "github.com/andelf/go-curl"
	"syscall"
	"unsafe"
	"fmt"
)

func FD_ZERO(set *syscall.FdSet) {
    s := unsafe.Pointer(set)
    C._FD_ZERO(s)
}

func FD_SET(sysfd int, set *syscall.FdSet) {
    s := unsafe.Pointer(set)
    fd := C.int(sysfd)
    C._FD_SET(fd, s)
}

func FD_ISSET(sysfd int, set *syscall.FdSet) bool {
    s := unsafe.Pointer(set)
    fd := C.int(sysfd)
    return C._FD_ISSET(fd, s) != 0
}

func main() {
    var (
            rset, wset, eset syscall.FdSet
            still_running, curl_timeout int = 0, 0
            err error
    )

	ch1 := curl.EasyInit()
	ch2 := curl.EasyInit()

	ch1.Setopt(curl.OPT_URL, "http://www.163.com")
	ch1.Setopt(curl.OPT_HEADER, 0)
	ch1.Setopt(curl.OPT_VERBOSE, true)
	ch2.Setopt(curl.OPT_URL, "http://www.baidu.com")
	ch2.Setopt(curl.OPT_HEADER, 0)
	ch2.Setopt(curl.OPT_VERBOSE, true)

	mh := curl.MultiInit()

	mh.AddHandle(ch1)
	mh.AddHandle(ch2)

	for {
        FD_ZERO(&rset)
        FD_ZERO(&wset)
        FD_ZERO(&eset)

        timeout := syscall.Timeval{Sec:1, Usec:0}
       	curl_timeout, err = mh.Timeout()
       	if err != nil {
       		fmt.Printf("Error multi_timeout: %s\n", err)
       	}
       	if curl_timeout >= 0 {
       		timeout.Sec = int64(curl_timeout / 1000)
       		if timeout.Sec > 1 {
       			timeout.Sec = 1
       		} else {
       			timeout.Usec = int64((curl_timeout % 1000)) * 1000
       		}
       	}

	max_fd, err := mh.Fdset(&rset, &wset, &eset)
        if err != nil {
            fmt.Printf("Error FDSET: %s\n", err)
        }

        _, err = syscall.Select(int(max_fd + 1), &rset, &wset, &eset, &timeout)
        if err != nil {
        	fmt.Printf("Error select: %s\n", err)
        } else {
        	still_running, err = mh.Perform()
        	if still_running > 0 {
        		fmt.Printf("Still running: %d\n", still_running)
        	} else {
        		break
        	}
        }
	}

}

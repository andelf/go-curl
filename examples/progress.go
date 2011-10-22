
package main

import (
	"../src/_obj/curl"
	"time"
	"fmt"
)


func write_data(ptr []byte, size uintptr, userdata interface{}) uintptr {
	// silent
	return size
}



func main() {
	curl.GlobalInit(curl.GLOBAL_ALL)

	// init the curl session
	easy := curl.EasyInit()
	defer easy.Cleanup()

	easy.Setopt(curl.OPT_URL, "http://curl.haxx.se/download/curl-7.22.0.tar.gz")

	easy.Setopt(curl.OPT_WRITEFUNCTION, write_data)

	easy.Setopt(curl.OPT_NOPROGRESS, false)

	easy.Setopt(curl.OPT_PROGRESSFUNCTION, func (_ interface{}, dltotal float64, dlnow float64, ultotal float64, ulnow float64) int {
		fmt.Printf("Download %.2f, Uploading %.2f\r", dlnow/dltotal, ulnow/ultotal)
		return 0
	})

	if err := easy.Perform(); err != nil {
		println("ERROR: ", err.String())
	}
}

// A sample program to show how to use PROGRESS callback to calculate
// downloading percentage and speed
package main

import (
	"fmt"
	curl "github.com/andelf/go-curl"
	"time"
)

func write_data(ptr []byte, userdata interface{}) bool {
	// make it ok, do nothing
	return true
}

func main() {
	curl.GlobalInit(curl.GLOBAL_ALL)

	// init the curl session
	easy := curl.EasyInit()
	defer easy.Cleanup()

	easy.Setopt(curl.OPT_URL, "http://curl.haxx.se/download/curl-7.22.0.tar.gz")

	easy.Setopt(curl.OPT_WRITEFUNCTION, write_data)

	easy.Setopt(curl.OPT_NOPROGRESS, false)

	started := int64(0)
	easy.Setopt(curl.OPT_PROGRESSFUNCTION, func(dltotal, dlnow, ultotal, ulnow float64, userdata interface{}) bool {
		// canceled when 50% finished
		if dlnow/dltotal > 0.5 {
			println("")
			// abort downloading
			return false
		}
		if started == 0 {
			started = time.Now().Unix()
		}
		fmt.Printf("Downloaded: %3.2f%%, Speed: %.1fKiB/s \r", dlnow/dltotal*100, dlnow/1000/float64((time.Now().Unix()-started)))
		return true
	})

	if err := easy.Perform(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
}

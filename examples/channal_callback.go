
package main

import (
	"../curl/_obj/curl"
	"time"
)

func write_data(ptr []byte, size uintptr, userdata interface{}) uintptr {
	ch, ok := userdata.(chan string)
	if ok {
		ch <- string(ptr)
		return size
	} else {
		println("ERROR!")
		return 0
	}
	return 0
}


func main() {
	curl.GlobalInit(curl.GLOBAL_ALL)

	// init the curl session
	easy := curl.EasyInit()
	defer easy.Cleanup()

	easy.Setopt(curl.OPT_URL, "http://cn.bing.com/")

	easy.Setopt(curl.OPT_WRITEFUNCTION, write_data)

	// make a chan
	ch := make(chan string, 100)
	go func(ch chan string) {
		for {
			data := <-ch
			println("Got data size=", len(data))
		}
	}(ch)

	easy.Setopt(curl.OPT_WRITEDATA, ch)

	if err := easy.Perform(); err != nil {
		println("ERROR: ", err.String())
	}

	time.Sleep(10000)			// wait gorotine
}

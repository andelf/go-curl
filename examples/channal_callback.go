package main

import (
	curl "github.com/andelf/go-curl"
	"time"
)

func write_data(ptr []byte, userdata interface{}) bool {
	ch, ok := userdata.(chan string)
	if ok {
		ch <- string(ptr)
		return true // ok
	} else {
		println("ERROR!")
		return false
	}
	return false
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
		println("ERROR: ", err.Error())
	}

	time.Sleep(10000) // wait gorotine
}

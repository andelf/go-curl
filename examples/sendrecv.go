
package main

import (
	"../src/_obj/curl"
	"time"
)

const POST_DATA = "a_test_data_only"

var sent = false

func main() {
	// init the curl session

	result := make([]byte, 1000)


	easy := curl.EasyInit()
	defer easy.Cleanup()

	easy.Setopt(curl.OPT_URL, "http://www.renren.com")


	easy.Setopt(curl.OPT_PORT, 80)
	easy.Setopt(curl.OPT_VERBOSE, true)
	easy.Setopt(curl.OPT_CONNECT_ONLY, true)





	easy.Setopt(curl.OPT_WRITEFUNCTION, nil)

	if err := easy.Perform(); err != nil {
		println("ERROR: ", err.String(), err)
	}

	easy.Send([]byte("HEAD / HTTP/1.0\r\nHost: www.renren.com\r\n\r\n"))


	time.Sleep(1000000000)			// wait gorotine
	if _, err := easy.Recv(result); err != nil {
		println("ERROR:", err.String())
	}
	println(string(result))

}

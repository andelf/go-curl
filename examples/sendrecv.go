package main

import (
	"fmt"
	curl "github.com/andelf/go-curl"
	"time"
)

const POST_DATA = "a_test_data_only"

var sent = false

func main() {
	// init the curl session

	easy := curl.EasyInit()
	defer easy.Cleanup()

	easy.Setopt(curl.OPT_URL, "http://www.renren.com")

	easy.Setopt(curl.OPT_PORT, 80)
	easy.Setopt(curl.OPT_VERBOSE, true)
	easy.Setopt(curl.OPT_CONNECT_ONLY, true)

	easy.Setopt(curl.OPT_WRITEFUNCTION, nil)

	if err := easy.Perform(); err != nil {
		println("ERROR: ", err.Error())
	}

	easy.Send([]byte("HEAD / HTTP/1.0\r\nHost: www.renren.com\r\n\r\n"))

	buf := make([]byte, 1000)
	time.Sleep(1000000000) // wait gorotine
	num, err := easy.Recv(buf)
	if err != nil {
		println("ERROR:", err.Error())
	}
	println("recv num = ", num)
	// NOTE: must use buf[:num]
	println(string(buf[:num]))

	fmt.Printf("got:\n%#v\n", string(buf[:num]))
}

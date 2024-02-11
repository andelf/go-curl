package main

import (
	curl "github.com/andelf/go-curl"
	"time"
)

const POST_DATA = "a_test_data_only"

var sent = false

func main() {
	// init the curl session
	easy := curl.EasyInit()
	defer easy.Cleanup()

	easy.Setopt(curl.OPT_URL, "http://www.google.com")

	easy.Setopt(curl.OPT_POST, true)
	easy.Setopt(curl.OPT_VERBOSE, true)

	easy.Setopt(curl.OPT_READFUNCTION,
		func(ptr []byte, userdata interface{}) int {
			// WARNING: never use append()
			if !sent {
				sent = true
				ret := copy(ptr, POST_DATA)
				return ret
			}
			return 0 // sent ok
		})

	// disable HTTP/1.1 Expect 100
	easy.Setopt(curl.OPT_HTTPHEADER, []string{"Expect:"})
	// must set
	easy.Setopt(curl.OPT_POSTFIELDSIZE, len(POST_DATA))

	if err := easy.Perform(); err != nil {
		println("ERROR: ", err.Error())
	}

	time.Sleep(10000) // wait gorotine
}

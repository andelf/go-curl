
package main

import (
	"../curl/_obj/curl"
	"os"
)

const (
	headerfilename = "head.out"
	bodyfilename = "body.out"
)

func write_data(ptr []byte, size uintptr, userdata interface{}) uintptr {
	//println("DEBUG(write_data): ", userdata)
	//println("DEBUG", userdata.(interface{}))
	fp := userdata.(* os.File)
	if writed, err := fp.Write(ptr); err == nil {
		return uintptr(writed)
	}

	return 0
}


func main() {
	curl.GlobalInit(curl.GLOBAL_ALL)

	// init the curl session
	easy := curl.EasyInit()
	defer easy.Cleanup()

	// set URL to get
	easy.Setopt(curl.OPT_URL, "http://cn.bing.com/")

	// no progress meter
	easy.Setopt(curl.OPT_NOPROGRESS, true)

	easy.Setopt(curl.OPT_WRITEFUNCTION, write_data)

	// write file
	fp, _ := os.OpenFile(bodyfilename, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0777)
	defer fp.Close()
	easy.Setopt(curl.OPT_WRITEDATA, fp)

	// easy.Setopt(curl.OPT_WRITEHEADER, 0)

	if err := easy.Perform(); err != nil {
		println("ERROR: ", err.String())
	}

}

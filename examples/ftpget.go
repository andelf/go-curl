package main

import (
	curl "github.com/andelf/go-curl"
	"os"
)

const filename = "README"

func main() {
	curl.GlobalInit(curl.GLOBAL_DEFAULT)
	defer curl.GlobalCleanup()
	easy := curl.EasyInit()
	defer easy.Cleanup()

	easy.Setopt(curl.OPT_URL, "ftp://ftp.gnu.org/README")

	// define our callback use lambda function
	easy.Setopt(curl.OPT_WRITEFUNCTION, func(ptr []byte, userdata interface{}) bool {
		file := userdata.(*os.File)
		if _, err := file.Write(ptr); err != nil {
			return false
		}
		return true
	})

	fp, _ := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	defer fp.Close() // defer close

	easy.Setopt(curl.OPT_WRITEDATA, fp)

	easy.Setopt(curl.OPT_VERBOSE, true)

	if err := easy.Perform(); err != nil {
		println("ERROR", err.Error())
	}
}

package main

import (
	"../curl/_obj/curl"
)

func main() {
	easy := curl.EasyInit()
	defer easy.Cleanup()
	if easy != nil {
		easy.Setopt(curl.OPT_URL, "http://www.google.com/")
		easy.Perform()
	}
}

package main

import (
	curl "github.com/andelf/go-curl"
)

func main() {
	easy := curl.EasyInit()
	defer easy.Cleanup()
	if easy != nil {
		easy.Setopt(curl.OPT_URL, "https://mail.google.com/")
		// skip_peer_verification
		easy.Setopt(curl.OPT_SSL_VERIFYPEER, false) // 0 is ok

		easy.Perform()
	}
}

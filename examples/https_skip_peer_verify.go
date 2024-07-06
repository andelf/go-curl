package main

import (
	curl "github.com/andelf/go-curl"
)

func main() {
	easy := curl.EasyInit()
	defer easy.Cleanup()
	if easy != nil {
		easy.Setopt(curl.OPT_URL, "https://mail.google.com/")

		// OPT_SSL_VERIFYPEER determines whether curl verifies the authenticity of the peer's certificate.
		// Do not disable OPT_SSL_VERIFYPEER unless you absolutely sure of the security implications.
		// https://curl.se/libcurl/c/CURLOPT_SSL_VERIFYPEER.html
		easy.Setopt(curl.OPT_SSL_VERIFYPEER, false) // 0 also means false

		easy.Perform()
	}
}

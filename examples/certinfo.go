package main

import (
	"fmt"
	"log"

	"github.com/andelf/go-curl"
)

func main() {
	const domain = "google.com"

	if err := curl.GlobalInit(curl.GLOBAL_ALL); err != nil {
		log.Println("failed to init libcurl")
		return
	}
	defer curl.GlobalCleanup()

	easy := curl.EasyInit()
	defer easy.Cleanup()

	if easy == nil {
		log.Println("failed to init easy curl")
		return
	}

	url := fmt.Sprintf("https://%s", domain)

	easy.Setopt(curl.OPT_URL, url)
	easy.Setopt(curl.OPT_SSL_VERIFYPEER, true)
	easy.Setopt(curl.OPT_SSL_VERIFYHOST, true)
	easy.Setopt(curl.OPT_TIMEOUT, 5)
	easy.Setopt(curl.OPT_CERTINFO, true)
	easy.Setopt(curl.OPT_NOPROGRESS, true)
	easy.Setopt(curl.OPT_NOBODY, true)

	if err := easy.Perform(); err != nil {
		log.Printf("failed to preform request: %v", err)
		return
	}

	res, err := easy.Getinfo(curl.INFO_CERTINFO)
	if err != nil {
		log.Printf("failed to get cert info: %v", err)
		return
	}

	info, ok := res.([]string)
	if !ok {
		log.Printf("unknown cert info result: %T", res)
		return
	}

	fmt.Printf("Fetched %d certificates!\n\n", len(info))
	for i, certInfo := range info {
		fmt.Printf("Certificate info %d:\n", i)
		fmt.Printf("%+v\n", certInfo)
		fmt.Println()
	}
}

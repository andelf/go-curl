package main

import (
	"fmt"
	"github.com/sorinescu/go-curl"
	"time"
)

func main() {
	ch1 := curl.EasyInit()
	ch2 := curl.EasyInit()

	ch1.Setopt(curl.OPT_URL, "http://www.163.com")
	ch1.Setopt(curl.OPT_HEADER, 0)
	ch1.Setopt(curl.OPT_VERBOSE, true)
	ch2.Setopt(curl.OPT_URL, "http://www.baidu.com")
	ch2.Setopt(curl.OPT_HEADER, 0)
	ch2.Setopt(curl.OPT_VERBOSE, true)

	mh := curl.MultiInit()

	mh.AddHandle(ch1)
	mh.AddHandle(ch2)

	var repeats int

	for {
		stillRunning, err := mh.Perform()
		if err != nil {
			fmt.Printf("Error perform: %s\n", err)
			break
		} else if stillRunning > 0 {
			fmt.Printf("Still running: %d\n", stillRunning)
		} else {
			break
		}

		timeoutMs := 1000
		curlTimeout, err := mh.Timeout()
		if err != nil {
			fmt.Printf("Error multi_timeout: %s\n", err)
		}
		if curlTimeout >= 0 {
			timeoutMs = curlTimeout
		}

		numFds, err := mh.Wait(timeoutMs)
		if err != nil {
			fmt.Printf("Error wait: %s\n", err)
			break
		}

		// 'numFds' being zero means either a timeout or no file descriptors to
		// wait for. Try timeout on first occurrence, then assume no file
		// descriptors, which means wait for 100 milliseconds to prevent spinning
		// in Perform + Wait.
		if numFds == 0 {
			repeats++ // count number of repeated zero numFds
			if repeats > 1 {
				time.Sleep(100 * time.Millisecond)
			}
		} else {
			repeats = 0
		}
	}
}

go-curl
=======

my golang libcurl(curl) binding.

 * under active development
 * by andelf

See more examples in ./examples/ directory~!

Current Development Statue
--------------------------

 * READ, WRITE, HEADER, PROGRESS function callback
 * a Multipart Form supports file uploading
 * Most curl_easy_setopt option
 * partly implement share & multi interface
 * new callback function prototype

How to Install
--------------

	$ goinstall github.com/andelf/go-curl/curl  # or newer
    $ go get -u github.com/andelf/go-curl/curl

Current Status
--------------

 * Linux x64
   * passed go release.r59 9022 (arch pacman)
   * passed go release.r60 9481 (arch pacman)
   * failed go release.r60.3 9516 (arch pacman, crosscall2: not defined)
   * passed go weekly.2011-10-18 10143 (gcc, hg repo)
   * passed go weekly.2011-10-18 10143 (clang, hg repo)
 * Windows x86
   * passed go release.r60.3 9516 (win7, mingw-gcc 4.5.2)

Sample Program
--------------

```go
package main

import (
    "fmt"
    "github.com/andelf/go-curl/curl"
)

func main() {
    easy := curl.EasyInit()
    defer easy.Cleanup()

    easy.Setopt(curl.OPT_URL, "http://www.baidu.com/")

    // make a callback function
    fooTest := func (buf []byte, userdata interface{}) bool {
        println("DEBUG: size=>", len(buf))
        println("DEBUG: content=>", string(buf))
        return true
    }

    easy.Setopt(curl.OPT_WRITEFUNCTION, fooTest)

    if err := easy.Perform(); err != nil {
        fmt.Printf("ERROR: %v\n", err)
    }
}
```

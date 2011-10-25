go-curl
=======

my golang libcurl(curl) binding.

 * already implemented go callback in cgo
 * under active development
 * by andelf

See more examples in ./examples/ directory~!


Current Status
--------------

 * Linux x64
   * passed go release.r59 9022 (arch pacman)
   * passed go release.r60 9481 (arch pacman)
   * failed go release.r60.3 9516 (arch pacman, crosscall2: not defined)
   * passed go weekly.2011-10-18 10143 (gcc, hg repo)
   * passed go weekly.2011-10-18 10143 (clang, hg repo)
 * Windows
   * passed go release.r60.3 9516 (win7, i686, mingw-gcc 4.5.2)

Sample Program
--------------

    package main

    import (
        "fmt"
        "../src/_obj/curl"
    )

    func main() {
        ret := curl.EasyInit()
        defer ret.Cleanup()

        if ret.Setopt(curl.OPT_URL, "http://www.baidu.com:8000/") != nil {
            println("set url ok!")
        }

        println("set resolve =>", ret.Setopt(curl.OPT_RESOLVE, []string{"www.baidu.com:8000:127.0.0.1",}) == nil)

        // make a callback function
        fooTest := func (buf []byte, size uintptr, userdata interface{}) uintptr {
            // NOTE: here size == len(buf)
            println("DEBUG: size=>", len(buf))
            println("DEBUG: params:", buf, size, userdata)
            println("DEBUG: =>", string(buf))
            return size				// must return size of byte
        }

        ret.Setopt(curl.OPT_WRITEFUNCTION, fooTest)

        if err := ret.Perform(); err != nil {
            fmt.Printf("error code -> %v\n", err)
        }
    }

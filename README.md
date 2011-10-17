my golang libcurl(curl) binding.

 * already implemented go callback in cgo
 * under active development
 * by andelf



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

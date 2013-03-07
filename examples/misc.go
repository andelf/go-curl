package main

import (
	"fmt"
	curl "github.com/andelf/go-curl"
	"os"
	"reflect"
)

const endl = "\n"

func main() {
	println("DEBUG chdir=>", os.Chdir("/sadf"))
	ret := curl.EasyInit()
	defer ret.Cleanup()
	print("init =>", ret, " ", reflect.TypeOf(ret).String(), endl)
	ret = ret.Duphandle()
	defer ret.Cleanup()

	print("dup =>", ret, " ", reflect.TypeOf(ret).String(), endl)
	print("global init =>", curl.GlobalInit(curl.GLOBAL_ALL), endl)
	print("version =>", curl.Version(), endl)
	// debug
	//print("set verbose =>", ret.Setopt(curl.OPT_VERBOSE, true), endl)

	//print("set header =>", ret.Setopt(curl.OPT_HEADER, true), endl)

	// auto calculate port
	// print("set port =>", ret.EasySetopt(curl.OPT_PORT, 6060), endl)
	fmt.Printf("XXXX debug setopt %#v \n", ret.Setopt(30000, 19).Error())

	print("set timeout =>", ret.Setopt(curl.OPT_TIMEOUT, 20), endl)

	//print("set post size =>", ret.Setopt(curl.OPT_POSTFIELDSIZE, 10), endl)
	if ret.Setopt(curl.OPT_URL, "http://www.baidu.com:8000/") != nil {
		println("set url ok!")
	}
	//print("set url =>", ret.Setopt(curl.OPT_URL, "http://commondatastorage.googleapis.com/chromium-browser-continuous/Linux_x64/104547/chrome-linux.zip"), endl)

	print("set user_agent =>", ret.Setopt(curl.OPT_USERAGENT, "go-curl v0.0.1") == nil, endl)
	// add to DNS cache
	print("set resolve =>", ret.Setopt(curl.OPT_RESOLVE, []string{"www.baidu.com:8000:127.0.0.1"}) == nil, endl)
	// ret.EasyReset()  clean seted

	// currently not finished!
	//
	fooTest := func(buf []byte, userdata interface{}) bool {
		// buf := ptr.([]byte)
		println("size=>", len(buf))
		println("DEBUG(in callback)", buf, userdata)
		println("data = >", string(buf))
		return true
	}

	ret.Setopt(curl.OPT_WRITEFUNCTION, fooTest) // curl.CallbackWriteFunction(fooTest))
	println("set opt!")
	// for test only

	code := ret.Perform()
	//	dump.Dump(code)
	fmt.Printf("code -> %v\n", code)

	println("================================")
	print("pause =>", ret.Pause(curl.PAUSE_ALL), endl)

	print("escape =>", ret.Escape("http://baidu.com/"), endl)
	print("unescape =>", ret.Unescape("http://baidu.com/-%00-%5c"), endl)

	print("unescape lenght =>", len(ret.Unescape("http://baidu.com/-%00-%5c")), endl)
	// print("version info data =>", curl.VersionInfo(1), endl)
	ver := curl.VersionInfo(curl.VERSION_NOW)
	fmt.Printf("VersionInfo: Age: %d, Version:%s, Host:%s, Features:%d, SslVer: %s, LibzV: %s, ssh: %s\n",
		ver.Age, ver.Version, ver.Host, ver.Features, ver.SslVersion, ver.LibzVersion, ver.LibsshVersion)

	print("Protocols:")
	for _, p := range ver.Protocols {
		print(p, ", ")
	}
	print(endl)
	println(curl.Getdate("20111002 15:05:58 +0800").String())
	ret.Getinfo(curl.INFO_EFFECTIVE_URL)
	ret.Getinfo(curl.INFO_RESPONSE_CODE)

	ret.Getinfo(curl.INFO_FILETIME)
	ret.Getinfo(curl.INFO_SSL_ENGINES)

	ret.Getinfo(curl.INFO_TOTAL_TIME)

	println("================================")

	// ret.Getinfo(curl.INFO_SSL_ENGINES)

	/*	mret := curl.MultiInit()
		mret.AddHandle(ret)			// works
		defer mret.Cleanup()
		if ok, handles := mret.Perform(); ok == curl.OK {
			fmt.Printf("ok=%s, handles=%d\n", ok, handles)
		} else {
			fmt.Printf("error calling multi\n")
		}
	*/
	println("================================")
	//println(curl.GlobalInit(curl.GLOBAL_SSL))
}

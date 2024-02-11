// 测试人人网登录, 并保存cookiejar
package main

import (
	"flag"
	curl "github.com/andelf/go-curl"
)

func main() {
	// init the curl session

	var username *string = flag.String("username", "test", "renren.com email")
	var password *string = flag.String("password", "test", "renren.com password")

	flag.Parse()

	easy := curl.EasyInit()
	defer easy.Cleanup()

	easy.Setopt(curl.OPT_URL, "http://m.renren.com/home.do")

	easy.Setopt(curl.OPT_PORT, 80)
	easy.Setopt(curl.OPT_VERBOSE, true)

	easy.Setopt(curl.OPT_COOKIEJAR, "./cookie.jar")

	// disable HTTP/1.1 Expect: 100-continue
	easy.Setopt(curl.OPT_HTTPHEADER, []string{"Expect:"})

	postdata := "email=" + *username + "&password=" + *password + "&login=" + easy.Escape("登录")
	easy.Setopt(curl.OPT_POSTFIELDS, postdata)

	if err := easy.Perform(); err != nil {
		println("ERROR: ", err.Error())
	}
}

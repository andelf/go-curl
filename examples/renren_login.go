// 测试人人网登录, 并保存cookiejar
package main

import (
	"github.com/andelf/go-curl/curl"
)

const POST_DATA = "a_test_data_only"

var sent = false

func main() {
	// init the curl session

	easy := curl.EasyInit()
	defer easy.Cleanup()

	easy.Setopt(curl.OPT_URL, "http://m.renren.com/home.do")

	easy.Setopt(curl.OPT_PORT, 80)
	easy.Setopt(curl.OPT_VERBOSE, true)

	easy.Setopt(curl.OPT_COOKIEJAR, "./cookie.jar")

	// disable HTTP/1.1 Expect: 100-continue
	easy.Setopt(curl.OPT_HTTPHEADER, []string{"Expect:"})

	// easy.Setopt(curl.OPT_HTTPPOST, form)
	easy.Setopt(curl.OPT_POSTFIELDS, "email=你的邮箱&password=你的密码&login="+easy.Escape("登录"))

	if err := easy.Perform(); err != nil {
		println("ERROR: ", err.Error())
	}
}

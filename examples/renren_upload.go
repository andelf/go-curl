// 人人网图片上传
package main

import (
	"fmt"
	curl "github.com/andelf/go-curl"
	"regexp"
	"time"
)

func getUploadUrl() string {
	page := ""
	easy := curl.EasyInit()
	defer easy.Cleanup()

	u := "http://3g.renren.com/album/wuploadphoto.do"
	easy.Setopt(curl.OPT_URL, u)
	easy.Setopt(curl.OPT_COOKIEFILE, "./cookie.jar")
	easy.Setopt(curl.OPT_COOKIEJAR, "./cookie.jar")
	easy.Setopt(curl.OPT_VERBOSE, true)
	easy.Setopt(curl.OPT_WRITEFUNCTION, func(ptr []byte, _ interface{}) bool {
		page += string(ptr)
		return true
	})
	easy.Perform()
	// extract url from
	// <form enctype="multipart/form-data" action="http://3g.renren.com/album/wuploadphoto.do?type=3&amp;sid=zv3tiXTZr6Cu1rj5dhgX_X"
	pattern, _ := regexp.Compile(`action="(.*?)"`)

	if matches := pattern.FindStringSubmatch(page); len(matches) == 2 {
		return matches[1]
	}
	return ""
}

func main() {
	// init the curl session

	easy := curl.EasyInit()
	defer easy.Cleanup()

	posturl := getUploadUrl()

	easy.Setopt(curl.OPT_URL, posturl)

	easy.Setopt(curl.OPT_PORT, 80)
	easy.Setopt(curl.OPT_VERBOSE, true)

	// save cookie and load cookie
	easy.Setopt(curl.OPT_COOKIEFILE, "./cookie.jar")
	easy.Setopt(curl.OPT_COOKIEJAR, "./cookie.jar")

	// disable HTTP/1.1 Expect: 100-continue
	easy.Setopt(curl.OPT_HTTPHEADER, []string{"Expect:"})

	form := curl.NewForm()
	form.Add("albumid", "452618633") // your album id
	form.AddFile("theFile", "./test.jpg")
	form.Add("description", "我就尝试下这段代码靠谱不。。截图下看看")
	form.Add("post", "上传照片")

	easy.Setopt(curl.OPT_HTTPPOST, form)

	// print upload progress
	easy.Setopt(curl.OPT_NOPROGRESS, false)
	easy.Setopt(curl.OPT_PROGRESSFUNCTION, func(dltotal, dlnow, ultotal, ulnow float64, _ interface{}) bool {
		fmt.Printf("Download %3.2f%%, Uploading %3.2f%%\r", dlnow/dltotal*100, ulnow/ultotal*100)
		return true
	})

	if err := easy.Perform(); err != nil {
		println("ERROR: ", err.Error())
	}

	time.Sleep(1000000000) // wait gorotine

}

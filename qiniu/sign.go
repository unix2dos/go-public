package main

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"strings"
	"time"
)

func main() {

	resUrl := "https://video-preview-fhyx.cloudfh.com/YUGAOPIAN_ZHUI-XIONG-SHI-JIU-NIAN_2.39_01.jpg"
	fmt.Println(resUrl)

	newUrl := getSignUrl(resUrl)
	fmt.Println(newUrl)
}

func getSignUrl(resUrl string) string {

	key := "058834992fabf38b05b94d8ba19b4199a09b49ab"

	resUri, err := url.Parse(resUrl)
	if err != nil {
		return resUrl
	}

	expireTime := fmt.Sprintf("%x", time.Now().Add(time.Second*10).Unix())
	rawStr := fmt.Sprintf("%s%s%s", key, resUri.Path, expireTime)

	m := md5.New()
	m.Write([]byte(rawStr))
	sign := fmt.Sprintf("%x", m.Sum(nil))

	var newUrl string
	if strings.Contains(resUrl, "?") {
		newUrl = fmt.Sprintf("%s&sign=%s&t=%s", resUrl, sign, expireTime)
	} else {
		newUrl = fmt.Sprintf("%s?sign=%s&t=%s", resUrl, sign, expireTime)
	}

	return newUrl
}

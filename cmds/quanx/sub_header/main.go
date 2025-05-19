package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cast"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

/*
AMY
Subscription-UserInfo
upload=2658137767; download=20245549223; total=53687091200; expire=1699574400


GlaDos
Subscription-Userinfo
upload=1; download=30725596593; total=214748364800; expire=1703217450

*/
func main() {
	router := gin.Default()
	router.GET("/shadow", shadow)
	router.GET("/glados", glados)
	router.TrustedPlatform = "X-CDN-IP"
	router.Run(":9091")
}

func glados(c *gin.Context) {
	var urlC, urlQ string
	var count int64

	name := c.Query("name")
	switch name {
	case "2": // 500G
		urlC = "https://update.glados-config.org/clash/130417/5138411/88714/glados.yaml"
		urlQ = "https://update.glados-config.com/quantumultx/130417/ba4b8edab08ddc16"
		count = gladosUsage("Cookie=enabled; Cookie.sig=lbtpENsrE0x6riM8PFTvoh9nepc; koa:sess=eyJ1c2VySWQiOjEzMDQxNywiX2V4cGlyZSI6MTY5NTM0MjMyMzk2NywiX21heEFnZSI6MjU5MjAwMDAwMDB9; koa:sess.sig=YRqd6u97FsK-OlRaku9nMkBSSmc")
	case "3": // 200G
		urlC = "https://update.glados-config.com/clash/122722/1f36602/153430/glados.yaml"
		urlQ = "https://update.glados-config.com/quantumultx/122722/a1d5677925852fca"
		count = gladosUsage("Cookie=enabled; Cookie.sig=lbtpENsrE0x6riM8PFTvoh9nepc; cf_clearance=QU0o65Dhf76G3ha1kMXvYVZmLo2ZBU.eRAs5S4pdcbk-1669422581-0-160; koa:sess=eyJ1c2VySWQiOjEyMjcyMiwiX2V4cGlyZSI6MTY5NTQzNjE0MjI2OSwiX21heEFnZSI6MjU5MjAwMDAwMDB9; koa:sess.sig=NoSBkwH-owFbBLAk_L7oXIkukr4")
	}
	fmt.Println("glados", urlC, urlQ)
	client := resty.New()

	respQ, _ := client.R().Get(urlQ)
	respQBody := respQ.String()
	respC, _ := client.R().Get(urlC)

	for k, v := range respQ.Header() {
		c.Header(k, v[0])
	}

	for k, v := range respC.Header() {
		if k == "Subscription-Userinfo" {
			sub := v[0]
			arr := strings.Split(sub, ";")
			for i := 0; i < len(arr); i++ {
				if strings.Contains(arr[i], "download=0") {
					arr[i] = fmt.Sprintf(" download=%d", count)
				}
			}
			sub2 := strings.Join(arr, ";")
			fmt.Println(k, sub)
			fmt.Println(k, sub2)
			c.Header(k, sub2)
		}
	}

	c.String(200, respQBody)
}

func gladosUsage(cookie string) (count int64) {
	url := "https://glados.one/api/user/usage"
	method := "GET"

	client := &http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authority", "glados.one")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")
	req.Header.Add("cookie", cookie)
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"107\", \"Chromium\";v=\"107\", \"Not=A?Brand\";v=\"24\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	type Temp struct {
		Code int       `json:"code"`
		Data [][]int64 `json:"data"`
	}

	var t Temp
	err = json.Unmarshal(body, &t)
	var usage int64
	for _, v := range t.Data {
		usage += v[1]
	}
	gb := float64(usage) / float64(1000000000)
	count = int64(gb * 1024 * 1024 * 1024)
	return
}

func shadow(c *gin.Context) {
	var url string
	name := c.Query("name")
	switch name {
	case "1":
		url = "https://s.trojanflare.com/subscription/quantumultx/b6402d33-484d-4a0b-8b3f-82424e1a55fd"
	case "2":
		url = "https://s.trojanflare.com/subscription/quantumultx/47aa849f-4228-4b84-ad87-aa5e4d9d09fe"
	}
	client := resty.New()
	resp, _ := client.R().Get(url)
	respBody := resp.String()
	fmt.Println("shadow", url)

	//服务有效至2025/11/11 剩78.85GB,
	var download float64
	var expire int64
	var day string
	{
		s := "服务有效至"
		index1 := strings.Index(respBody, s)
		index2 := strings.Index(respBody, " 剩")
		str := respBody[index1+len(s) : index2]
		t, _ := time.Parse("2006/01/02", str)
		day = t.Format("2006-01-02 15:04:05")
		expire = t.Unix()
	}
	{
		s := " 剩"
		index1 := strings.Index(respBody, s)
		index2 := strings.Index(respBody, "GB,")
		str := respBody[index1+len(s) : index2]
		download = 100 - cast.ToFloat64(str)
	}

	for k, v := range resp.Header() {
		c.Header(k, v[0])
	}
	sub := fmt.Sprintf("upload=0; download=%d; total=%d; expire=%d", int(download*1024*1024*1024), 100*1024*1024*1024, expire)
	c.Header("Subscription-Userinfo", sub)
	fmt.Println("Subscription-Userinfo", sub)
	fmt.Println("download", download, "expire", day)
	c.String(200, respBody)
}

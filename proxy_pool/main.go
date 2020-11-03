package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {

	proxy, err := url.Parse("http://101.200.199.4:3128")
	if err != nil {
		log.Fatal(err)
	}

	httpClient := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}

	res, err := httpClient.Get("https://www.liuvv.com")
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Println(err)
		return
	}
	c, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(c))
}

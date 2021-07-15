package main

import (
	"encoding/json"
	"fmt"
)

type Info struct {
	A string `json:"a"`
	B int    `json:"b"`
	C string `json:"c"`
}

func main() {
	info := &Info{A: "str1", B: 100, C: "str2"}
	o := struct {
		Info
		interface
	}{}

	outputJson(info)
}

func outputJson(res interface{}) {
	data, _ := json.Marshal(res)
	fmt.Println(string(data))
}

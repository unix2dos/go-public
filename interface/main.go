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

type AA struct {
	*Info
	string
}

func main() {
	info := &Info{A: "str1", B: 100, C: "str2"}

	// 这里来了两个测试的动态结构
	vars1 := `{"d":"123"}`
	//vars2 := `{"e":"123", "f":200}`

	//outPuts := make(map[string]interface{})
	//
	//respByte, _ := json.Marshal(info)
	//_ = json.Unmarshal(respByte, &outPuts)
	//
	//_ = json.Unmarshal([]byte(vars1), &outPuts)
	//_ = json.Unmarshal([]byte(vars2), &outPuts)

	o := struct {
		*Info
		interface{}
	}{info, vars1}

	outputJson(o)
}

// 理解成给前端的json
func outputJson(res interface{}) {
	data, _ := json.Marshal(res)
	fmt.Println(string(data))
}

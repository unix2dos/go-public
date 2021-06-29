package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Fixed struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	o := struct {
		Fixed
		Time string `json:"time"`
	}{
		Fixed: Fixed{Name: "levon", Age: 9},
		Time:  time.Now().String(),
	}
	outputJson(o)
}

// 理解成给前端的json
func outputJson(res interface{}) {
	data, _ := json.Marshal(res)
	fmt.Println(string(data))
}

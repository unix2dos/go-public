package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	str := "A"
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(str)))

	str1 := "A00A"
	fmt.Println(base64.URLEncoding.EncodeToString([]byte(str1)))

	b, _ := base64.StdEncoding.DecodeString("QQQQQQ==")
	fmt.Println(b)
}

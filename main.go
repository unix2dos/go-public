package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ReverseWord("never too later to learn"))
}

// 输入： never too later to learn
// 输出： learn to later too never

func ReverseWord(str string) string {
	arr := strings.Fields(str)
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return strings.Join(arr, " ")
}

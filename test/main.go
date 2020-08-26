package main

import "fmt"

func main() {
	s := []int{5, 4, 3, 2, 1}
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Println(s[i])
	}
}

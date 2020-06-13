package main

import (
	"fmt"
)

func main() {
	func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("a")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("b")
	}()

	panic("异常信息")

	fmt.Println("c")
}

package main

import "fmt"

type ReadWriter interface {
	~string | ~[]rune

	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

func main() {
	b := BytesReadWriter("str")
	b.Write([]byte("hello"))

	//var i ReadWriter  // 不行，因为不能定义变量
	//i = b
	//i.Write([]byte("world"))
}

// 类型 StringReadWriter 实现了接口 Readwriter
type StringReadWriter string

func (s StringReadWriter) Read(p []byte) (n int, err error) {
	return
}
func (s StringReadWriter) Write(p []byte) (n int, err error) {
	fmt.Println("111", string(p))
	return
}

// 类型BytesReadWriter 没有实现接口 Readwriter, 既不是string也是不[]rune
type BytesReadWriter []byte

func (s BytesReadWriter) Read(p []byte) (n int, err error) {
	return
}
func (s BytesReadWriter) Write(p []byte) (n int, err error) {
	return
}

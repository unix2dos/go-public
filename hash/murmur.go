package main

import (
	"crypto/md5"
	"fmt"

	"github.com/spaolacci/murmur3"
)

func main() {
	a := murmur3.New32()
	a.Write([]byte("1234"))
	fmt.Println(fmt.Sprintf("%x", a.Sum32()))

	b := murmur3.New64()
	b.Write([]byte("1234"))
	fmt.Println(fmt.Sprintf("%x", b.Sum64()))

	fmt.Println(fmt.Sprintf("%x", murmur3.Sum64([]byte("1234"))))

	c := murmur3.New128()
	c.Write([]byte("1234"))
	fmt.Println(c.Sum128())

	fmt.Println(fmt.Sprintf("%x", md5.Sum([]byte("1234"))))

}

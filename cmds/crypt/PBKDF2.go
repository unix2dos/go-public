package main

import (
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

func main() {

	passwd := "levonfly"
	salt := "salt"

	res1 := pbkdf2.Key([]byte(passwd), []byte(salt), 10, 10, sha256.New)
	fmt.Println(string(res1)) //'J!85|LU@

	// 加密后一样
	res2 := pbkdf2.Key([]byte(passwd), []byte(salt), 10, 20, sha256.New)
	fmt.Println(string(res2)) //'J!85|LU@

}

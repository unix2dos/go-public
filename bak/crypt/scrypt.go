package main

import (
	"fmt"

	"golang.org/x/crypto/scrypt"
)

func main() {

	passwd := "levonfly"
	salt := "salt"

	res1, _ := scrypt.Key([]byte(passwd), []byte(salt), 1<<15, 8, 1, 32)
	fmt.Println(string(res1)) //TCoi[DRt;IALuw}

	res2, _ := scrypt.Key([]byte(passwd), []byte(salt), 1<<15, 8, 1, 32)
	fmt.Println(string(res2)) //TCoi[DRt;IALuw}
}

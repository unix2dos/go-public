package main

import (
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

func main() {

	passwd := "levonfly"
	salt := "salt"

	res1 := argon2.IDKey([]byte(passwd), []byte(salt), 3, 32, 4, 300)
	fmt.Println(base64.StdEncoding.EncodeToString(res1)) //uEZgAbCSfDyd8VAMbcmSSZKpH/TQ9hh9VsblPFGuDjM
	fmt.Println(string(res1))                            //uEZgAbCSfDyd8VAMbcmSSZKpH/TQ9hh9VsblPFGuDjM

	res2 := argon2.IDKey([]byte(passwd), []byte(salt), 3, 32, 4, 30)
	fmt.Println(base64.StdEncoding.EncodeToString(res2)) //uEZgAbCSfDyd8VAMbcmSSZKpH/TQ9hh9VsblPFGuDjM
	fmt.Println(string(res2))
}

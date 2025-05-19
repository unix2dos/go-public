package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	passwd := "levonfly"

	res1, _ := bcrypt.GenerateFromPassword([]byte(passwd), 10)
	fmt.Println(string(res1)) //$2a$10$Y85p96ZRD1Sa5iU7M/ngku9MIFNkmAwEI38FvPT9dj628E8hPOU0K

	// 第二次加密的结果和第一次不一样
	res2, _ := bcrypt.GenerateFromPassword([]byte(passwd), 10)
	fmt.Println(string(res2)) //$2a$10$7xUWgmWB3te5OipBYx4aheUFz7dCcj7JLIpQW6D/Me1R4qljEIFy2

	err1 := bcrypt.CompareHashAndPassword(res1, []byte(passwd))
	fmt.Println(err1) //nil

	err2 := bcrypt.CompareHashAndPassword(res1, []byte("random"))
	fmt.Println(err2) //crypto/bcrypt: hashedPassword is not the hash of the given password
}

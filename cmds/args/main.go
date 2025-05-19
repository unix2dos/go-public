package main

import "fmt"

type ArgFunc func() (key string, value string)

func BuildArgFunc(key string, value string) ArgFunc {
	return func() (k string, v string) {
		k = key
		v = value
		return
	}
}

func ExecUser(name string, age int, funcArgs ...ArgFunc) {
	fmt.Println("name:", name, "age:", age)
	for _, funcArg := range funcArgs {
		k, v := funcArg()
		fmt.Println(k, v)
	}
}

func main() {
	ExecUser("levonfly", 9)
	ExecUser("levonfly", 9, BuildArgFunc("email", "levonfly@gmail.com"))
	ExecUser("levonfly", 9, BuildArgFunc("email", "levonfly@gmail.com"), BuildArgFunc("sex", "man"))
}

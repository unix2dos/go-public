package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	name := os.Args[1]
	fi, err := os.Open(name)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		str := string(a)
		str = strings.ToLower(str)
		str = str[0:len(str)-3]


		max := letterMaxCount(str)
		if max > 2{
			fmt.Println(str)
		}
	}
}


func shouwei(str string) bool{
	if str[0] == str[len(str)-1]{
		return true
	}
	return false
}


func letterMaxCount(str string) int{
	max := 0
	for i := 'a'; i <= 'z'; i++ {
		a := string(i)
		n := strings.Count(str, a)
		if max < n {
			max = n
		}
	}
	return max
}

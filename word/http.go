package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fi, err := os.Open("http.txt")
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
		arr := strings.Split(string(a), ".com")
		for _, v := range arr {
			if len(v) > 7 {
				continue
			}
			fmt.Println(strings.ToLower(v))
		}
	}
}

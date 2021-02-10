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
	head := uint8('a')
	count := 0
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		str := string(a)
		str = strings.ToLower(str)
		str = str[0:len(str)-4]

		if str[0] == str[len(str)-1]{
			//fmt.Println("首尾", str)
		}

		max := 0
		maxA := ""
		for i := 'a'; i <= 'z'; i++ {
			a := string(i)
			n := strings.Count(str, a)
			if max < n {
				max = n
				maxA = a
			}
		}



		if max > 4{
			//if head == str[0] {
			//	count++
			//}else{
			//	fmt.Println("--------------------------",string(head),count)
			//	head = str[0]
			//	count=1
			//}
			_ = head
			_= count
			fmt.Println(str,max,maxA)
		}
	}
}

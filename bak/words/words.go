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
	var all []string
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		str := string(a)
		str = strings.ToLower(str)

		arr := strings.Split(str, " ")
		all = append(all, arr...)
	}

	allmap := map[string]string{}
	for _, v := range all {
		if v != "" {
			allmap[v] = v
			fmt.Println(v)
		}
	}
	fmt.Println("all", len(allmap))

	//{
	//	name := os.Args[2]
	//	fi, err := os.Open(name)
	//	if err != nil {
	//		fmt.Printf("Error: %s\n", err)
	//		return
	//	}
	//	defer fi.Close()
	//
	//	br := bufio.NewReader(fi)
	//	for {
	//		a, _, c := br.ReadLine()
	//		if c == io.EOF {
	//			break
	//		}
	//		str := string(a)
	//		str = strings.ToLower(str)
	//
	//		//if _, ok := allmap[str]; ok {
	//		//	fmt.Println(str)
	//		//}
	//
	//		for _, vv := range all {
	//			//fmt.Println("111", vv, str)
	//			if str == vv {
	//				fmt.Println(str, vv)
	//			}
	//		}
	//	}
	//}
}

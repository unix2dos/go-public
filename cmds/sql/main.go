package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cast"
)

func main() {
	heat := Heat()
	im := IM()
	black := Black()
	user := UserLog()

	fmt.Println("heat", len(heat))
	fmt.Println("im", len(im))
	fmt.Println("user", len(user))
	fmt.Println("black", len(black))

	intersect := GetIntersectionInt64(GetIntersectionInt64(heat, im), user)
	fmt.Println("intersect", len(intersect))

	res := Block(intersect, black)
	res = Unique(res)
	fmt.Println("res", len(res))
	WriteMaptoFile(res, "res.txt")
}

func GetIntersectionInt64(s1 []int64, s2 []int64) (data []int64) {
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v1 == v2 {
				data = append(data, v1)
			}
		}
	}
	return data
}

func UserLog() []int64 {
	data := make([]int64, 0)
	fi, err := os.Open("./user_log.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return []int64{}
	}
	defer fi.Close()

	mapUser := make(map[int64]int, 0)

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		arr := strings.Split(string(a), "\t")
		userId := cast.ToInt64(arr[0])
		//data := arr[1]
		mapUser[userId]++
	}

	for k, v := range mapUser {
		if v >= 10 {
			data = append(data, k)
		}
	}
	return data
}

func WriteMaptoFile(res []int64, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, v := range res {
		lineStr := fmt.Sprintf("%d", v)
		fmt.Fprintln(w, lineStr)
	}
	return w.Flush()
}

func Unique(input []int64) []int64 {
	u := make([]int64, 0, len(input))
	m := make(map[int64]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}

func Block(origin []int64, block []int64) []int64 {

	blockMap := make(map[int64]bool, 0)
	for _, v := range block {
		blockMap[v] = true
	}
	res := make([]int64, 0)
	for _, v := range origin {
		if _, ok := blockMap[v]; !ok {
			res = append(res, v)
		}
	}
	return res
}

func Black() []int64 {
	data := make([]int64, 0)
	fi, err := os.Open("./report_user.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return []int64{}
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		data = append(data, cast.ToInt64(string(a)))
	}
	return data
}

func IM() []int64 {
	data := make([]int64, 0)
	fi, err := os.Open("./im_sum_user.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return []int64{}
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		data = append(data, cast.ToInt64(string(a)))
	}
	return data
}

func Heat() []int64 {
	heat := make([]int64, 0)
	fi, err := os.Open("./heat.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return []int64{}
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		heat = append(heat, cast.ToInt64(string(a)))
	}
	return heat
}

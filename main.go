package main

import (
	"fmt"
)

func main() {
	//var wg sync.WaitGroup
	ch := make(chan int)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range arr {
		//wg.Add(1)
		go func(v int) {
			//defer wg.Done()
			ch <- v
		}(v)

		fmt.Println(v)

	}

	//wg.Wait()
}

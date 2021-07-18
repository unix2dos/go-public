package main

import "fmt"

func main() {
	arr := []int{21, 2, 1, 60, 32}
	Print(arr) //21      2       1       60      32
	QuickSort(arr, 0, len(arr)-1)
	Print(arr) //1       2       21      32      60
}

func QuickSort(arr []int, left, right int) {
	if left < right {
		mid := partition(arr, left, right)
		QuickSort(arr, left, mid-1)
		QuickSort(arr, mid+1, right)
	}
}

func partition(arr []int, left, right int) int {

	value := arr[left] // 两人心中所想的数字
	start := left      // 记录所想数字的位置，相遇了要交换

	for left != right { // 两人不相遇，就循环

		for left < right && arr[right] >= value { //牛郎先走，大于等于就走，小于就停
			right--
		}
		for left < right && arr[left] <= value { //织女后走，小于等于就走，大于就停
			left++
		}

		if left < right { //双方停止后打电话交换脚下的数据
			arr[left], arr[right] = arr[right], arr[left]
		}
		//接着循环，直到两人相遇
	}

	// 相遇了，把两人脚下数据和心中数字交换，下面 left 和 right 是相等的
	arr[start], arr[left] = arr[left], arr[start]

	return left
}

func Print(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d\t", v)
	}
	fmt.Print("\n")
}

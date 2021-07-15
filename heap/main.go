package main

import "fmt"

type List struct {
	Val  int
	Next *List
}

func main() {
	list4 := &List{Val: 4, Next: nil}
	list3 := &List{Val: 3, Next: list4}
	list2 := &List{Val: 2, Next: list3}
	list1 := &List{Val: 1, Next: list2}
	printList(list1)
	listOK := swapList(list1)
	printList(listOK)
}

func printList(head *List) {
	tmp := head
	for tmp != nil {
		fmt.Printf("%d\t", tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println("")
}

func swapList(head *List) *List {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	ok := swapList(head.Next.Next)
	head.Next.Next = head
	head.Next = ok
	return next
}
